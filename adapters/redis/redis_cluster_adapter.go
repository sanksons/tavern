package redis

import "github.com/go-redis/redis"

const SLOTS_COUNT = 16384

type RedisClusterConfig struct {
	Addrs []string
}

type RedisCluster struct {
	redisbase
}

func (this *RedisCluster) Initialize(config RedisClusterConfig) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: config.Addrs,
	})
	this.Client = client
}
