package redis

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

const (
	statsKey  = "stats"
	separator = "-"
)

var client *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

type TopRequest struct {
	Hits  int64  `json:"hits"`
	Int1  int64  `json:"int1"`
	Int2  int64  `json:"int2"`
	Limit int64  `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// increments given request in redis
func IncrHitRequest(req *http.Request) error {
	query := req.URL.Query()
	params := []string{
		query["int1"][0],
		query["int2"][0],
		query["limit"][0],
		query["str1"][0],
		query["str2"][0],
	}
	_, err := client.ZIncrBy(statsKey, 1, strings.Join(params, separator)).Result()
	return err
}

func GetTopRequest() (*TopRequest, error) {
	nbOfKey, err := client.Exists(statsKey).Result()
	if err != nil {
		return nil, err
	}
	if nbOfKey == 0 {
		return nil, fmt.Errorf(`key "%s" does not exists`, statsKey)
	}

	topRequests, err := client.ZRevRangeWithScores(statsKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	strMember, ok := topRequests[0].Member.(string)
	if !ok {
		return nil, fmt.Errorf(`something went wrong while retrieving top request params in redis`)
	}
	reqParams := strings.Split(strMember, separator)
	int1, err := strconv.ParseInt(reqParams[0], 10, 64)
	if err != nil {
		return nil, err
	}
	int2, err := strconv.ParseInt(reqParams[0], 10, 64)
	if err != nil {
		return nil, err
	}
	limit, err := strconv.ParseInt(reqParams[0], 10, 64)
	if err != nil {
		return nil, err
	}

	return &TopRequest{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Hits:  int64(topRequests[0].Score),
		Str1:  reqParams[3],
		Str2:  reqParams[4],
	}, nil

}