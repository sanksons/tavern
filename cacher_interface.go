package tavern

import (
	"fmt"

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
	Get(utils.CacheKey) ([]byte, error)
	Set(utils.CacheItem) error
	MGet(...utils.CacheKey) (map[utils.CacheKey][]byte, error)
	MSet(...utils.CacheItem) (map[utils.CacheKey]bool, error)
	Destroy(...utils.CacheKey) (map[utils.CacheKey]bool, error)
}

func Initialize(adapter string, config interface{}) (CacheAdapter, error) {
	switch adapter {
	case ADAPTER_TYPE_LOCAL:
		c, ok := config.(local.LocalAdapterConfig)
		if !ok {
			return nil, fmt.Errorf("Expected local.LocalAdapterConfig, Got: %T", config)
		}
		return local.Initialize(c), nil
	case ADAPTER_TYPE_REDIS_SIMPLE:
		c, ok := config.(redis.RedisSimpleConfig)
		if !ok {
			return nil, fmt.Errorf("Expected redis.RedisSimpleConfig, Got: %T", config)
		}
		return redis.InitializeRedisSimple(c), nil
	case ADAPTER_TYPE_REDIS_CLUSTER:
		c, ok := config.(redis.RedisClusterConfig)
		if !ok {
			return nil, fmt.Errorf("Expected redis.RedisClusterConfig, Got: %T", config)
		}
		return redis.InitializeRedisCluster(c), nil
	default:
		return nil, fmt.Errorf("Not a valid adapter type supplied")
	}
}
