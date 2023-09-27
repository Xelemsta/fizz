package monitoring

import (
	"fizz/restapi/operations/monitoring"

	"github.com/go-openapi/runtime/middleware"
)

type getMonPing struct{}

func NewGetMonPingHandler() monitoring.GetMonPingHandler {
	return &getMonPing{}
}

// Handle implements GET /mon/ping
func (impl *getMonPing) Handle(params monitoring.GetMonPingParams) middleware.Responder {
	ok := "OK"

	return monitoring.NewGetMonPingOK().WithPayload(&monitoring.GetMonPingOKBody{Status: &ok})
}
