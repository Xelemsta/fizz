package metrics

import (
	"fizz/restapi/operations/metrics"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type getMetrics struct{}

func NewGetMetricsHandler() metrics.GetMetricsHandler {
	return &getMetrics{}
}

// Handle implements GET /metrics.
// Used by prometheus instance.
func (impl *getMetrics) Handle(params metrics.GetMetricsParams) middleware.Responder {
	return metricsResponder(func(w http.ResponseWriter, _ runtime.Producer) {
		promhttp.Handler().ServeHTTP(w, params.HTTPRequest)
	})
}

type metricsResponder func(http.ResponseWriter, runtime.Producer)

func (mr metricsResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	mr(w, p)
}
