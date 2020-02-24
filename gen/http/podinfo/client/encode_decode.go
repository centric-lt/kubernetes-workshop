// Code generated by goa v3.0.3, DO NOT EDIT.
//
// podinfo HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/centric-lt/k8s-101/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	podinfo "github.com/centric-lt/k8s-101/gen/podinfo"
	podinfoviews "github.com/centric-lt/k8s-101/gen/podinfo/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "podinfo" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetPodinfoPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("podinfo", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the podinfo
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("podinfo", "get", err)
			}
			p := NewGetPodinforesultOK(&body)
			view := "default"
			vres := &podinfoviews.Podinforesult{p, view}
			if err = podinfoviews.ValidatePodinforesult(vres); err != nil {
				return nil, goahttp.ErrValidationError("podinfo", "get", err)
			}
			res := podinfo.NewPodinforesult(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("podinfo", "get", resp.StatusCode, string(body))
		}
	}
}
