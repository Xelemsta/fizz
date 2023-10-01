package datastore_test

import (
	"fizz/internal/datastore"
	"net/http"
	"testing"
	"time"

	"github.com/maxatome/go-testdeep/td"
)

func TestMapIncrHit(t *testing.T) {
	cases := []struct {
		label         string
		url           string
		expectedError error
	}{
		{
			label:         "ok",
			url:           "/v1/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz",
			expectedError: nil,
		},
	}

	for _, c := range cases {
		f := func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, c.url, nil)
			td.CmpNoError(t, err)
			td.Cmp(t, datastore.GetMapStats().IncrHitRequest(req), c.expectedError)
		}
		t.Run(c.label, f)
	}
}

func TestMapTopRequest(t *testing.T) {
	cases := []struct {
		label              string
		before             func()
		expectedError      error
		expectedTopRequest *datastore.TopRequest
	}{
		{
			label: "happy halloween three hits",
			before: func() {
				backend, err := datastore.GetBackend(string(datastore.MapBackendName))
				td.CmpNoError(t, err)
				td.CmpNotNil(t, backend)

				req, err := http.NewRequest(http.MethodGet, "/v1/fizzbuzz?int1=5&int2=7&limit=100&str1=happy&str2=halloween", nil)
				td.CmpNoError(t, err)
				for _, req := range []*http.Request{req, req, req} {
					err = backend.IncrHitRequest(req)
					td.CmpNoError(t, err)
				}

				req, err = http.NewRequest(http.MethodGet, "/v1/fizzbuzz?int1=8&int2=10&limit=100&str1=fizz&str2=buzz", nil)
				td.CmpNoError(t, err)
				err = backend.IncrHitRequest(req)
				td.CmpNoError(t, err)

			},
			expectedError: nil,
			expectedTopRequest: &datastore.TopRequest{
				Hits:  3,
				Int1:  5,
				Int2:  7,
				Limit: 100,
				Str1:  "happy",
				Str2:  "halloween",
			},
		},
		{
			label: "fizz buzz four hits",
			before: func() {
				backend, err := datastore.GetBackend(string(datastore.MapBackendName))
				td.CmpNoError(t, err)
				td.CmpNotNil(t, backend)

				req, err := http.NewRequest(http.MethodGet, "/v1/fizzbuzz?int1=8&int2=10&limit=100&str1=fizz&str2=buzz", nil)
				td.CmpNoError(t, err)
				for _, req := range []*http.Request{req, req, req} {
					err = backend.IncrHitRequest(req)
					td.CmpNoError(t, err)
				}

				req, err = http.NewRequest(http.MethodGet, "/v1/fizzbuzz?int1=5&int2=7&limit=100&str1=alone&str2=request", nil)
				td.CmpNoError(t, err)
				err = backend.IncrHitRequest(req)
				td.CmpNoError(t, err)
			},
			expectedError: nil,
			expectedTopRequest: &datastore.TopRequest{
				Hits:  4,
				Int1:  8,
				Int2:  10,
				Limit: 100,
				Str1:  "fizz",
				Str2:  "buzz",
			},
		},
	}

	for _, c := range cases {
		f := func(t *testing.T) {
			if c.before != nil {
				c.before()
				time.Sleep(500 * time.Millisecond)
			}
			topRequest, err := datastore.GetMapStats().GetTopRequest()
			if err != nil {
				td.Cmp(t, err.Error(), c.expectedError.Error())
			} else {
				td.Cmp(t, err, c.expectedError)
			}
			td.Cmp(t, topRequest, c.expectedTopRequest)
		}
		t.Run(c.label, f)
	}
}
