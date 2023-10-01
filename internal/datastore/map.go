package datastore

import (
	"net/http"
	"strconv"
	"strings"
)

var mapStats MapStats = make(map[string]int64)

type MapStats map[string]int64

func GetMapStats() MapStats {
	return mapStats
}

// IncrHitRequest increments given request hits in a in memory map
func (m MapStats) IncrHitRequest(req *http.Request) error {
	query := req.URL.Query()
	key := generateMemberFromQueryArgs(
		query["int1"][0],
		query["int2"][0],
		query["limit"][0],
		query["str1"][0],
		query["str2"][0],
	)

	m[key] += 1
	return nil
}

// GetTopRequest retrieves top count of api requests (including query args)
func (m MapStats) GetTopRequest() (*TopRequest, error) {
	var topCount int64
	var topKey string
	for key, count := range m {
		if count > topCount {
			topCount = count
			topKey = key
		}
	}

	queryArgs := strings.Split(topKey, queryArgsSeparator)
	int1, err := strconv.ParseInt(queryArgs[0], 10, 64)
	if err != nil {
		return nil, err
	}
	int2, err := strconv.ParseInt(queryArgs[1], 10, 64)
	if err != nil {
		return nil, err
	}
	limit, err := strconv.ParseInt(queryArgs[2], 10, 64)
	if err != nil {
		return nil, err
	}

	return &TopRequest{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Hits:  topCount,
		Str1:  queryArgs[3],
		Str2:  queryArgs[4],
	}, nil
}
