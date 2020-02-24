package cmd

import (
	"context"
	"fmt"
	"github.com/centric-lt/k8s-101/controllers"
	podinfosvr "github.com/centric-lt/k8s-101/gen/http/podinfo/server"
	"github.com/centric-lt/k8s-101/gen/podinfo"
	"github.com/spf13/cobra"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
)

var httpServerCmd = &cobra.Command{
	Use:   "server",
	Short: "serves an application by starting an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		server(debug)
	},
}

func server(debug bool) {
	// Define command line flags, add any other flag required to configure the
	// service.
	logger := log.New(os.Stderr, "[pods] ", log.Ltime)
	podinfoSvc := controllers.NewPodinfo(logger)
	podinfoEndpoints := podinfo.NewEndpoints(podinfoSvc)

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	u, err := url.Parse("http://0.0.0.0:8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid URL: %s", err.Error())
		os.Exit(1)
	}
	handleHTTPServer(ctx, u, podinfoEndpoints, &wg, errc, logger, debug)
	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)
	// Send cancellation signal to the goroutines.
	cancel()
	wg.Wait()
	logger.Println("exited")
}

func handleHTTPServer(ctx context.Context, u *url.URL, podinfoEndpoints *podinfo.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {
	// Setup goa log adapter.
	var (
		adapter       = middleware.NewLogger(logger)
		dec           = goahttp.RequestDecoder
		enc           = goahttp.ResponseEncoder
		mux           = goahttp.NewMuxer()
		podinfoServer *podinfosvr.Server
	)
	{
		eh := errorHandler(logger)
		podinfoServer = podinfosvr.New(podinfoEndpoints, mux, dec, enc, eh)
	}
	// Configure the mux.
	podinfosvr.Mount(mux, podinfoServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if debug {
			handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
		}
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range podinfoServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
