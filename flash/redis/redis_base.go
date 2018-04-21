package redis

import (
	"github.com/go-redis/redis"
	"github.com/sanksons/flashcache/flash/common"
)

type Redis struct {
	Client RedisClient
}

func (this *Redis) Set(i common.CacheItem) error {
	err := this.Client.Set(string(i.Key), i.Value, i.Expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (this *Redis) Get(key string) ([]byte, error) {
	str, err := this.Client.Get(key).Result()
	if err != nil && err == redis.Nil {
		return nil, common.KeyNotExists
	}
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func (this *Redis) MGet(a ...string) (map[string][]byte, error) {
	sliceI, err := this.Client.MGet(a...).Result()
	if err != nil {
		return nil, err
	}
	m := make(map[string][]byte)
	for k, v := range sliceI {
		valstr, ok := v.(string)
		if !ok {
			//m[a[k]] = nil
			continue //skip
		}
		m[a[k]] = []byte(valstr)
	}

	return m, nil
}
