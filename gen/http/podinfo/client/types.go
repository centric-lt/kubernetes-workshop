// Code generated by goa v3.0.3, DO NOT EDIT.
//
// podinfo HTTP client types
//
// Command:
// $ goa gen github.com/centric-lt/k8s-101/design

package client

import (
	podinfoviews "github.com/centric-lt/k8s-101/gen/podinfo/views"
	goa "goa.design/goa/v3/pkg"
)

// GetOKResponseBody is the type of the "podinfo" service "get" endpoint HTTP
// response body.
type GetOKResponseBody struct {
	// POD ip address
	IP *string `form:"ip,omitempty" json:"ip,omitempty" xml:"ip,omitempty"`
	// POD hostname
	Hostname *string `form:"hostname,omitempty" json:"hostname,omitempty" xml:"hostname,omitempty"`
}

// GetInternalServerErrorResponseBody is used to define fields on response body
// types.
type GetInternalServerErrorResponseBody struct {
	// POD ip address
	IP *string `form:"ip,omitempty" json:"ip,omitempty" xml:"ip,omitempty"`
	// POD hostname
	Hostname *string `form:"hostname,omitempty" json:"hostname,omitempty" xml:"hostname,omitempty"`
}

// NewGetPodinforesultOK builds a "podinfo" service "get" endpoint result from
// a HTTP "OK" response.
func NewGetPodinforesultOK(body *GetOKResponseBody) *podinfoviews.PodinforesultView {
	v := &podinfoviews.PodinforesultView{
		IP:       body.IP,
		Hostname: body.Hostname,
	}
	return v
}

// ValidateGetInternalServerErrorResponseBody runs the validations defined on
// GetInternal Server ErrorResponseBody
func ValidateGetInternalServerErrorResponseBody(body *GetInternalServerErrorResponseBody) (err error) {
	if body.IP == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ip", "body"))
	}
	if body.Hostname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("hostname", "body"))
	}
	return
}
