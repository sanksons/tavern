package redis

import "github.com/go-redis/redis"

type RedisSimpleConfig struct {
	Addr string
}

type RedisSimple struct {
	redisbase
}

func (this *RedisSimple) Initialize(config RedisSimpleConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	this.Client = client
}
