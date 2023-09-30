package stats

import (
	"fizz/internal/redis"
	"fizz/internal/stats"
	"fizz/models"
	operation "fizz/restapi/operations/stats"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/juju/errors"
)

type getStats struct{}

func NewGetStatsHandler() operation.GetV1StatsHandler {
	return &getStats{}
}

// Handle implements GET /v1/metrics
func (impl *getStats) Handle(params operation.GetV1StatsParams) middleware.Responder {
	topRequest, err := stats.GetTopRequest(redis.GetClient())
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, errors.BadRequest) {
			status = http.StatusBadRequest
		}
		return operation.NewGetV1StatsDefault(status).WithPayload(&models.Error{
			Code:    int64(status),
			Message: err.Error(),
		})
	}

	return operation.NewGetV1StatsOK().WithPayload(&models.MostUsedRequest{
		Hits:  topRequest.Hits,
		Int1:  &topRequest.Int1,
		Int2:  &topRequest.Int2,
		Limit: &topRequest.Limit,
		Str1:  &topRequest.Str1,
		Str2:  &topRequest.Str2,
	})
}
