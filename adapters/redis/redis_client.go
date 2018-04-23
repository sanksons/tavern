package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var _ RedisClient = (*RedisSimpleClient)(nil)
var _ RedisClient = (*RedisClusterClient)(nil)

type RedisClient interface {
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	Get(string) *redis.StringCmd
	MGet(...string) *redis.SliceCmd
	MSet(...interface{}) *redis.StatusCmd
	Del(...string) *redis.IntCmd
}

type RedisSimpleClient struct {
	*redis.Client
}

type RedisClusterClient struct {
	*redis.ClusterClient
}
