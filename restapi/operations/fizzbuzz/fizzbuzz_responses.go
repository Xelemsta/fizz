// Code generated by go-swagger; DO NOT EDIT.

package fizzbuzz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"fizz/models"
)

// FizzbuzzOKCode is the HTTP code returned for type FizzbuzzOK
const FizzbuzzOKCode int = 200

/*
FizzbuzzOK fizz buzz string

swagger:response fizzbuzzOK
*/
type FizzbuzzOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewFizzbuzzOK creates FizzbuzzOK with default headers values
func NewFizzbuzzOK() *FizzbuzzOK {

	return &FizzbuzzOK{}
}

// WithPayload adds the payload to the fizzbuzz o k response
func (o *FizzbuzzOK) WithPayload(payload string) *FizzbuzzOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fizzbuzz o k response
func (o *FizzbuzzOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FizzbuzzOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
FizzbuzzDefault Error handling the request

swagger:response fizzbuzzDefault
*/
type FizzbuzzDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFizzbuzzDefault creates FizzbuzzDefault with default headers values
func NewFizzbuzzDefault(code int) *FizzbuzzDefault {
	if code <= 0 {
		code = 500
	}

	return &FizzbuzzDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the fizzbuzz default response
func (o *FizzbuzzDefault) WithStatusCode(code int) *FizzbuzzDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the fizzbuzz default response
func (o *FizzbuzzDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the fizzbuzz default response
func (o *FizzbuzzDefault) WithPayload(payload *models.Error) *FizzbuzzDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fizzbuzz default response
func (o *FizzbuzzDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FizzbuzzDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}