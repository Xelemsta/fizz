package stats

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

const (
	key       = "stats"
	separator = "-"
)

type TopRequest struct {
	Hits  int64  `json:"hits"`
	Int1  int64  `json:"int1"`
	Int2  int64  `json:"int2"`
	Limit int64  `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func generateMemberFromRequest(int1, int2, limit, str1, str2 string) string {
	return strings.Join([]string{int1, int2, limit, str1, str2}, separator)
}

// IncrHitRequest increments given request hits in redis
func IncrHitRequest(req *http.Request, client *redis.Client) error {
	query := req.URL.Query()
	member := generateMemberFromRequest(
		query["int1"][0],
		query["int2"][0],
		query["limit"][0],
		query["str1"][0],
		query["str2"][0],
	)

	_, err := client.ZIncrBy(key, 1, member).Result()
	return err
}

// GetTopRequest retrieves top count of api requests (with query args)
func GetTopRequest(client *redis.Client) (*TopRequest, error) {
	nbOfKey, err := client.Exists(key).Result()
	if err != nil {
		return nil, err
	}
	if nbOfKey == 0 {
		return nil, fmt.Errorf(`key %s does not exists`, key)
	}

	topRequests, err := client.ZRevRangeWithScores(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	strMember, ok := topRequests[0].Member.(string)
	if !ok {
		return nil, fmt.Errorf(`unexpected top request format (expected 'string')`)
	}
	reqParams := strings.Split(strMember, separator)
	int1, err := strconv.ParseInt(reqParams[0], 10, 64)
	if err != nil {
		return nil, err
	}
	int2, err := strconv.ParseInt(reqParams[1], 10, 64)
	if err != nil {
		return nil, err
	}
	limit, err := strconv.ParseInt(reqParams[2], 10, 64)
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
