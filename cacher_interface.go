package tavern

import (
	"github.com/sanksons/tavern/adapters/local"
	"github.com/sanksons/tavern/adapters/redis"
	"github.com/sanksons/tavern/utils"
)

const ADAPTER_TYPE_LOCAL = "local"
const ADAPTER_TYPE_REDIS_SIMPLE = "redis-simple"
const ADAPTER_TYPE_REDIS_CLUSTER = "redis-cluster"

var _ CacheAdapter = (*redis.RedisSimple)(nil)
var _ CacheAdapter = (*redis.RedisCluster)(nil)
var _ CacheAdapter = (*local.Local)(nil)

type CacheAdapter interface {
	Get(string) ([]byte, error)
	Set(utils.CacheItem) error
	MGet(...string) (map[string][]byte, error)
	MSet(...utils.CacheItem) (map[string]bool, error)
	Destroy(...string) (map[string]bool, error)
}
