package stats_test

import (
	internalRedis "fizz/internal/redis"
	"fizz/internal/stats"
	"fizz/testutils"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestIncrRequest(t *testing.T) {
	miniredis, err := miniredis.Run()
	td.CmpNoError(t, err)
	client := redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})

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
			td.Cmp(t, stats.IncrHitRequest(req, client), c.expectedError)
		}
		t.Run(c.label, f)
	}
}

func TestTopRequest(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))
	miniredis, err := miniredis.Run()
	td.CmpNoError(t, err)
	client := redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})

	internalRedis.SetClient(redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	}))

	cases := []struct {
		label              string
		before             func()
		expectedError      error
		expectedTopRequest *stats.TopRequest
	}{
		{
			label:              "key does not exist",
			before:             nil,
			expectedError:      fmt.Errorf(`key stats does not exists`),
			expectedTopRequest: nil,
		},
		{
			label: "happy halloween three hits",
			before: func() {
				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 8, "int2": 10, "limit": 100, "str1": "fizz", "str2": "buzz"},
				).CmpStatus(http.StatusOK).OrDumpResponse()

				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 5, "int2": 7, "limit": 100, "str1": "happy", "str2": "halloween"},
				).CmpStatus(http.StatusOK).OrDumpResponse()

				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 5, "int2": 7, "limit": 100, "str1": "happy", "str2": "halloween"},
				).CmpStatus(http.StatusOK).OrDumpResponse()

				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 5, "int2": 7, "limit": 100, "str1": "happy", "str2": "halloween"},
				).CmpStatus(http.StatusOK).OrDumpResponse()
			},
			expectedError: nil,
			expectedTopRequest: &stats.TopRequest{
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
				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 8, "int2": 10, "limit": 100, "str1": "fizz", "str2": "buzz"},
				).CmpStatus(http.StatusOK).OrDumpResponse()

				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 8, "int2": 10, "limit": 100, "str1": "fizz", "str2": "buzz"},
				).CmpStatus(http.StatusOK).OrDumpResponse()

				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 8, "int2": 10, "limit": 100, "str1": "fizz", "str2": "buzz"},
				).CmpStatus(http.StatusOK).OrDumpResponse()

				ta.Get(
					"/v1/fizzbuzz",
					tdhttp.Q{"int1": 5, "int2": 7, "limit": 100, "str1": "alone", "str2": "request"},
				).CmpStatus(http.StatusOK).OrDumpResponse()
			},
			expectedError: nil,
			expectedTopRequest: &stats.TopRequest{
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
			topRequest, err := stats.GetTopRequest(client)
			td.Cmp(t, err, c.expectedError)
			td.Cmp(t, topRequest, c.expectedTopRequest)
		}
		t.Run(c.label, f)
	}
}
