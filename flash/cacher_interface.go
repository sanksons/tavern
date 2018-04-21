package flash

import (
	"github.com/sanksons/flashcache/flash/common"
	"github.com/sanksons/flashcache/flash/redis"
)

const ADAPTER_TYPE_LOCAL = "local"
const ADAPTER_TYPE_REDIS_SIMPLE = "redis-simple"
const ADAPTER_TYPE_REDIS_CLUSTER = "redis-cluster"

var _ CacheAdapter = (*redis.RedisSimple)(nil)
var _ CacheAdapter = (*redis.RedisCluster)(nil)
var _ CacheAdapter = (*Local)(nil)

type CacheAdapter interface {
	Get(string) ([]byte, error)
	Set(common.CacheItem) error
	MGet(...string) (map[string][]byte, error)
	//MSet(map[string][]byte) map[string]bool
	//Destroy(string) error
}
