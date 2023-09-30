package redis

import "github.com/go-redis/redis"

var client *redis.Client

func SetClient(c *redis.Client) {
	client = c
}

func GetClient() *redis.Client {
	return client
}
