package datastore

import (
	"fmt"
	"net/http"
	"strings"
)

type BackendName string

const (
	RedisBackendName BackendName = "redis"
	MapBackendName   BackendName = "map"
	separator        string      = "-"
)

// TopRequest stores top request made
// by the api clients (according to params).
type TopRequest struct {
	Hits  int64  `json:"hits"`
	Int1  int64  `json:"int1"`
	Int2  int64  `json:"int2"`
	Limit int64  `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// Backend is a generic interface for a datastore backend
type Backend interface {
	IncrHitRequest(req *http.Request) error
	GetTopRequest() (*TopRequest, error)
}

// GetBackend returns provided "backend" (if implemented)
func GetBackend(backend string) (Backend, error) {
	switch backend {
	case string(RedisBackendName):
		return GetRedisClient(), nil
	case string(MapBackendName):
		return GetMapStats(), nil
	default:
		return nil, fmt.Errorf(`backend %s not implemented`, backend)
	}
}

// "5-8-100-fizz-buzz"
func generateMemberFromQueryParams(int1, int2, limit, str1, str2 string) string {
	return strings.Join([]string{int1, int2, limit, str1, str2}, separator)
}
