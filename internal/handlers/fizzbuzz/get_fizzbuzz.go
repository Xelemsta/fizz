package fizzbuzz

import (
	"fizz/restapi/operations/fizzbuzz"

	"github.com/go-openapi/runtime/middleware"
)

type fizzBuzzImpl struct{}

func NewFizzBuzzHandler() fizzbuzz.FizzbuzzHandler {
	return &fizzBuzzImpl{}
}

// Handle implements GET /v1/fizzbuzz.
func (impl *fizzBuzzImpl) Handle(params fizzbuzz.FizzbuzzParams) middleware.Responder {
	return fizzbuzz.NewFizzbuzzOK().WithPayload("test")
}
