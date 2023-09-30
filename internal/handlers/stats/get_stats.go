package stats

import (
	"net/http"

	"fizz/internal/datastore"
	"fizz/internal/transform"
	"fizz/models"
	"fizz/restapi/operations/stats"

	"github.com/go-openapi/runtime/middleware"
	"github.com/juju/errors"
)

type getStats struct{}

func NewGetStatsHandler() stats.GetV1StatsHandler {
	return &getStats{}
}

// Handle implements GET /v1/metrics
func (impl *getStats) Handle(params stats.GetV1StatsParams) middleware.Responder {
	backend, err := datastore.GetBackend(string(datastore.RedisBackendName))
	if err != nil {
		return stats.NewGetV1StatsDefault(http.StatusInternalServerError).WithPayload(&models.Error{
			Code:    int64(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	topRequest, err := backend.GetTopRequest()
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, errors.BadRequest) {
			status = http.StatusBadRequest
		}
		return stats.NewGetV1StatsDefault(status).WithPayload(&models.Error{
			Code:    int64(status),
			Message: err.Error(),
		})
	}

	return stats.NewGetV1StatsOK().WithPayload(transform.TopRequest(topRequest))
}
