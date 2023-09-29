package stats

import (
	"fizz/internal/redis"
	"fizz/models"
	"fizz/restapi/operations/stats"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

type getStats struct{}

func NewGetStatsHandler() stats.GetV1StatsHandler {
	return &getStats{}
}

// Handle implements GET /v1/metrics
func (impl *getStats) Handle(params stats.GetV1StatsParams) middleware.Responder {
	topRequest, err := redis.GetTopRequest()
	if err != nil {
		return stats.NewGetV1StatsDefault(http.StatusInternalServerError)
	}
	return stats.NewGetV1StatsOK().WithPayload(&models.MostUsedRequest{
		Hits:  topRequest.Hits,
		Int1:  &topRequest.Int1,
		Int2:  &topRequest.Int2,
		Limit: &topRequest.Limit,
		Str1:  &topRequest.Str1,
		Str2:  &topRequest.Str2,
	})
}
