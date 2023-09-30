package stats_test

import (
	internalRedis "fizz/internal/redis"
	"fizz/testutils"
	"net/http"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestStats(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))
	miniredis, err := miniredis.Run()
	td.CmpNoError(t, err)
	internalRedis.SetClient(redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	}))

	cases := []struct {
		label              string
		before             func()
		expectedHttpStatus int
		expectedJSONBody   any
	}{
		{
			label:              "no key",
			before:             nil,
			expectedHttpStatus: http.StatusInternalServerError,
			expectedJSONBody:   td.JSON(`{"code":500,"message":"key stats does not exists"}`),
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
			expectedHttpStatus: http.StatusOK,
			expectedJSONBody:   td.JSON(`{"hits":3,"int1":5,"int2":7,"limit":100,"str1":"happy","str2":"halloween"}`),
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
			expectedHttpStatus: http.StatusOK,
			expectedJSONBody:   td.JSON(`{"hits":4,"int1":8,"int2":10,"limit":100,"str1":"fizz","str2":"buzz"}`),
		},
	}

	for _, c := range cases {
		f := func(t *testing.T) {
			if c.before != nil {
				c.before()
				time.Sleep(500 * time.Millisecond)
			}
			ta.Get("/v1/stats").CmpStatus(c.expectedHttpStatus).CmpJSONBody(c.expectedJSONBody).OrDumpResponse()
		}
		t.Run(c.label, f)
	}
}
