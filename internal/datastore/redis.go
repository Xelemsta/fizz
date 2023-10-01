package datastore

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/juju/errors"
)

var redisClient *RedisClient

const (
	key = "stats"
)

type RedisClient struct {
	*redis.Client
}

func GetRedisClient() *RedisClient {
	if redisClient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})
		redisClient = &RedisClient{client}
	}
	return redisClient
}

// Mostly usefull for tests purpose
func SetRedisClient(c *RedisClient) {
	redisClient = c
}

// IncrHitRequest increments given request hits in redis
// using sorted set data type.
func (c *RedisClient) IncrHitRequest(req *http.Request) error {
	if redisClient == nil {
		return fmt.Errorf(`backend not initialized yet`)
	}

	query := req.URL.Query()
	member := generateMemberFromQueryArgs(
		query["int1"][0],
		query["int2"][0],
		query["limit"][0],
		query["str1"][0],
		query["str2"][0],
	)

	_, err := redisClient.ZIncrBy(key, 1, member).Result()
	return err
}

// GetTopRequest retrieves top count of api requests (including query args).
// If two api requests have the same count, they are lexicographically ordered.
func (c *RedisClient) GetTopRequest() (*TopRequest, error) {
	if c == nil {
		return nil, fmt.Errorf(`please provide a non nil redis client`)
	}
	nbOfKey, err := c.Exists(key).Result()
	if err != nil {
		return nil, err
	}
	if nbOfKey == 0 {
		return nil, errors.BadRequestf(`you need to perform at least one request before being able to retrieve top request`)
	}

	topRequests, err := c.ZRevRangeWithScores(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	strMember, ok := topRequests[0].Member.(string)
	if !ok {
		return nil, fmt.Errorf(`unexpected top request format (expected 'string')`)
	}
	queryArgs := strings.Split(strMember, queryArgsSeparator)
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
		Hits:  int64(topRequests[0].Score),
		Str1:  queryArgs[3],
		Str2:  queryArgs[4],
	}, nil
}
