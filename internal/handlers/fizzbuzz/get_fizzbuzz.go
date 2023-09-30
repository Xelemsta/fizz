package fizzbuzz

import (
	"fizz/internal/redis"
	"fizz/internal/stats"
	"fizz/models"
	"fizz/restapi/operations/fizzbuzz"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

const (
	min       int64  = 1
	max       int64  = 100
	separator string = ","
)

type fizzBuzzImpl struct{}

func NewFizzBuzzHandler() fizzbuzz.FizzbuzzHandler {
	return &fizzBuzzImpl{}
}

// Handle implements GET /v1/fizzbuzz.
func (impl *fizzBuzzImpl) Handle(params fizzbuzz.FizzbuzzParams) middleware.Responder {
	if params.Limit < min || params.Limit > max {
		return fizzbuzz.NewFizzbuzzDefault(http.StatusBadRequest).WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf(`limit must be between %d and %d, got %d`, min, max, params.Limit),
		})
	}

	// increments the counter of given request
	go func() {
		err := stats.IncrHitRequest(params.HTTPRequest, redis.GetClient())
		if err != nil {
			logrus.WithError(err).Warnf(`error while incrementing hit request`)
		}
	}()

	output := ""
	for i := min; i <= params.Limit; i++ {
		sep := separator
		if i == min {
			sep = ""
		}
		output += fizzBuzz(i, params.Int1, params.Int2, sep, params.Str1, params.Str2)
	}

	return fizzbuzz.NewFizzbuzzOK().WithPayload(&models.FizzBuzzResponse{
		Output: output,
	})
}

func fizzBuzz(i, int1, int2 int64, sep, str1, str2 string) string {
	// number is multiple of both int1 and int2
	if i%int1 == 0 && i%int2 == 0 {
		return sep + str1 + str2
	}
	// number is multiple of int1 only
	if i%int1 == 0 {
		return sep + str1
	}
	// number is multiple of int2 only
	if i%int2 == 0 {
		return sep + str2
	}
	// number is NOT a multiple of both int1 and int2
	return sep + strconv.FormatInt(i, 10)
}
