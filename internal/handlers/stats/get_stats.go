package stats

import (
	"fizz/models"
	"fizz/restapi/operations/stats"

	"github.com/go-openapi/runtime/middleware"
)

type getStats struct{}

func NewGetStatsHandler() stats.GetV1StatsHandler {
	return &getStats{}
}

// Handle implements GET /v1/metrics
func (impl *getStats) Handle(params stats.GetV1StatsParams) middleware.Responder {
	return stats.NewGetV1StatsOK().WithPayload(&models.MostUsedRequest{})
}
