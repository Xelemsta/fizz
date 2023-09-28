// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetV1StatsParams creates a new GetV1StatsParams object
//
// There are no default values defined in the spec.
func NewGetV1StatsParams() GetV1StatsParams {

	return GetV1StatsParams{}
}

// GetV1StatsParams contains all the bound params for the get v1 stats operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV1Stats
type GetV1StatsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV1StatsParams() beforehand.
func (o *GetV1StatsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
