package datastore_test

import (
	"fizz/internal/datastore"
	"fmt"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestDatastore(t *testing.T) {
	cases := []struct {
		label         string
		backendName   string
		expectedError error
	}{
		{
			label:         "not implemented",
			backendName:   "myBackend",
			expectedError: fmt.Errorf(`backend myBackend not implemented`),
		},
		{
			label:         "ok redis",
			backendName:   string(datastore.RedisBackendName),
			expectedError: nil,
		},
		{
			label:         "ok map",
			backendName:   string(datastore.MapBackendName),
			expectedError: nil,
		},
	}

	for _, c := range cases {
		f := func(t *testing.T) {
			backend, err := datastore.GetBackend(c.backendName)
			td.Cmp(t, err, c.expectedError)
			if err != nil {
				td.CmpNil(t, backend)
			}
		}
		t.Run(c.label, f)
	}
}
