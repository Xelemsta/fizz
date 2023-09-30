package fizzbuzz_test

import (
	"fizz/testutils"
	"net/http"
	"testing"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestFizzBuzz(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))
	cases := []struct {
		label              string
		queryArgs          tdhttp.Q
		int1               int64
		int2               int64
		limit              int64
		str1               string
		str2               string
		expectedHttpStatus int
		expectedJSONBody   any
	}{
		{
			label:              "no query args",
			queryArgs:          tdhttp.Q{},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":602,"message":"int1 in query is required"}`),
		},
		{
			label: "not positive int1",
			queryArgs: tdhttp.Q{
				"int1":  -1,
				"int2":  5,
				"limit": 50,
				"str1":  "fizz",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":609,"message":"int1 in query should be greater than or equal to 1"}`),
		},
		{
			label: "not positive int2",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  -1,
				"limit": 50,
				"str1":  "fizz",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":609,"message":"int2 in query should be greater than or equal to 1"}`),
		},
		{
			label: "zero int1",
			queryArgs: tdhttp.Q{
				"int1":  0,
				"int2":  5,
				"limit": 50,
				"str1":  "fizz",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":609,"message":"int1 in query should be greater than or equal to 1"}`),
		},
		{
			label: "zero int2",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  0,
				"limit": 50,
				"str1":  "fizz",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":609,"message":"int2 in query should be greater than or equal to 1"}`),
		},
		{
			label: "empty str1 and str2",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 50,
				"str1":  "",
				"str2":  "",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":602,"message":"str1 in query is required"}`),
		},
		{
			label: "empty str1",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 50,
				"str1":  "",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":602,"message":"str1 in query is required"}`),
		},
		{
			label: "empty str2",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 50,
				"str1":  "fizz",
				"str2":  "",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":602,"message":"str2 in query is required"}`),
		},
		{
			label: "invalid limit bottom",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 0,
				"str1":  "fizz",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":400,"message":"limit must be between 1 and 100, got 0"}`),
		},
		{
			label: "invalid limit top",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 101,
				"str1":  "fizz",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":400,"message":"limit must be between 1 and 100, got 101"}`),
		},
		{
			label: "forbidden char str1",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 30,
				"str1":  "fizz-",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":400,"message":"- is a forbidden char"}`),
		},
		{
			label: "forbidden char str2",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 30,
				"str1":  "fizz",
				"str2":  "buzz-",
			},
			expectedHttpStatus: http.StatusBadRequest,
			expectedJSONBody:   td.JSON(`{"code":400,"message":"- is a forbidden char"}`),
		},
		{
			label: "ok",
			queryArgs: tdhttp.Q{
				"int1":  3,
				"int2":  5,
				"limit": 30,
				"str1":  "fizz",
				"str2":  "buzz",
			},
			expectedHttpStatus: http.StatusOK,
			expectedJSONBody:   td.JSON(`{"output": "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz"}`),
		},
		{
			label: "another ok",
			queryArgs: tdhttp.Q{
				"int1":  5,
				"int2":  7,
				"limit": 100,
				"str1":  "happy",
				"str2":  "halloween",
			},
			expectedHttpStatus: http.StatusOK,
			expectedJSONBody:   td.JSON(`{"output":"1,2,3,4,happy,6,halloween,8,9,happy,11,12,13,halloween,happy,16,17,18,19,happy,halloween,22,23,24,happy,26,27,halloween,29,happy,31,32,33,34,happyhalloween,36,37,38,39,happy,41,halloween,43,44,happy,46,47,48,halloween,happy,51,52,53,54,happy,halloween,57,58,59,happy,61,62,halloween,64,happy,66,67,68,69,happyhalloween,71,72,73,74,happy,76,halloween,78,79,happy,81,82,83,halloween,happy,86,87,88,89,happy,halloween,92,93,94,happy,96,97,halloween,99,happy"}`),
		},
	}

	for _, c := range cases {
		f := func(t *testing.T) {
			ta.Get(
				"/v1/fizzbuzz",
				c.queryArgs,
			).CmpStatus(c.expectedHttpStatus).CmpJSONBody(c.expectedJSONBody).OrDumpResponse()
		}
		t.Run(c.label, f)
	}
}
