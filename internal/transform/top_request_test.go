package transform_test

import (
	"fizz/internal/datastore"
	"fizz/internal/transform"
	"fizz/models"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestTopRequest(t *testing.T) {
	hits := int64(1)
	int1 := int64(2)
	int2 := int64(3)
	limit := int64(4)
	str1 := "Str1"
	str2 := "Str2"
	cases := []struct {
		label               string
		topRequest          *datastore.TopRequest
		expectedApiResponse *models.MostUsedRequest
	}{
		{
			label: "ok",
			topRequest: &datastore.TopRequest{
				Hits:  1,
				Int1:  2,
				Int2:  3,
				Limit: 4,
				Str1:  "Str1",
				Str2:  "Str2",
			},
			expectedApiResponse: &models.MostUsedRequest{
				Hits:  &hits,
				Int1:  &int1,
				Int2:  &int2,
				Limit: &limit,
				Str1:  &str1,
				Str2:  &str2,
			},
		},
	}

	for _, c := range cases {
		f := func(t *testing.T) {
			td.Cmp(t, transform.TopRequest(c.topRequest), c.expectedApiResponse)
		}
		t.Run(c.label, f)
	}
}
