// Code generated by goa v3.0.3, DO NOT EDIT.
//
// podinfo HTTP server
//
// Command:
// $ goa gen github.com/centric-lt/k8s-101/design

package server

import (
	"context"
	"net/http"
	"path"
	"strings"

	podinfo "github.com/centric-lt/k8s-101/gen/podinfo"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the podinfo service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Get    http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the podinfo service endpoints.
func New(
	e *podinfo.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Get", "GET", "/pod"},
			{"static/index.html", "GET", "/"},
			{"static/", "GET", "/ui"},
		},
		Get: NewGetHandler(e.Get, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "podinfo" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Get = m(s.Get)
}

// Mount configures the mux to serve the podinfo endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetHandler(mux, h.Get)
	MountStaticIndexHTML(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	}))
	MountStatic(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upath := path.Clean(r.URL.Path)
		rpath := upath
		if strings.HasPrefix(upath, "/ui") {
			rpath = upath[3:]
		}
		http.ServeFile(w, r, path.Join("static/", rpath))
	}))
}

// MountGetHandler configures the mux to serve the "podinfo" service "get"
// endpoint.
func MountGetHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/pod", f)
}

// NewGetHandler creates a HTTP handler which loads the HTTP request and calls
// the "podinfo" service "get" endpoint.
func NewGetHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeGetResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get")
		ctx = context.WithValue(ctx, goa.ServiceKey, "podinfo")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountStaticIndexHTML configures the mux to serve GET request made to "/".
func MountStaticIndexHTML(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/", h.ServeHTTP)
}

// MountStatic configures the mux to serve GET request made to "/ui".
func MountStatic(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/ui/", h.ServeHTTP)
	mux.Handle("GET", "/ui/*path", h.ServeHTTP)
}
