package testutils

import (
	"fizz/restapi"
	"fizz/restapi/operations"
	"net/http"
	"testing"

	openapierrors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
)

func InitAPI(t testing.TB) http.Handler {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		t.Fatalf("Cannot load spec: %s", err)
	}
	api := operations.NewFizzBuzzAPIAPI(swaggerSpec)

	server := restapi.NewServer(api)
	server.ConfigureAPI()

	openapierrors.DefaultHTTPCode = http.StatusBadRequest

	return server.GetHandler()
}
