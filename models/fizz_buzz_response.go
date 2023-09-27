// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FizzBuzzResponse fizz buzz response
//
// swagger:model FizzBuzzResponse
type FizzBuzzResponse struct {

	// output
	Output string `json:"output,omitempty"`
}

// Validate validates this fizz buzz response
func (m *FizzBuzzResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fizz buzz response based on context it is used
func (m *FizzBuzzResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FizzBuzzResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FizzBuzzResponse) UnmarshalBinary(b []byte) error {
	var res FizzBuzzResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
