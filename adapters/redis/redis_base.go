package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/sanksons/tavern/utils"
)

// A common implementation for redis cluster and simple redis adapters.
// Anything common b/w the 2 is kept here.
// Anything specific is kept within the corresponding adapters itself.
type redisbase struct {
	Client RedisClient
}

// Implementation of Set method of CacheAdapter
func (this *redisbase) Set(i utils.CacheItem) error {
	err := this.Client.Set(string(i.Key), i.Value, i.Expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

//Implementation of Get method of CacheAdapter
func (this *redisbase) Get(key string) ([]byte, error) {
	str, err := this.Client.Get(key).Result()
	if err != nil && err == redis.Nil {
		return nil, utils.KeyNotExists
	}
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

//Implementation of MGet method of CacheAdapter
func (this *redisbase) MGet(a ...string) (map[string][]byte, error) {
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

//Implementation of MSet method of CacheAdapter
//@todo: should i use chunking or it is wise to use it at client side?
func (this *redisbase) MSet(items ...utils.CacheItem) (map[string]bool, error) {
	if len(items) == 0 {
		return nil, nil
	}
	results := make(map[string]bool)
	expiration := make(map[string]time.Duration)
	//transform
	transformedData := make([]interface{}, 0, len(items)*2)
	for _, i := range items {
		transformedData = append(transformedData, string(i.Key), i.Value)
		results[string(i.Key)] = true
		// set expiration
		if i.Expiration > 0 {
			expiration[string(i.Key)] = i.Expiration
		}
	}
	pipeline := this.Client.Pipeline()
	pipeline.MSet(transformedData...)
	for k, v := range expiration {
		pipeline.Expire(k, v)
	}
	_, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	return results, nil
}

//Implementation of Destroy method of CacheAdapter
func (this *redisbase) Destroy(keys ...string) (map[string]bool, error) {
	if len(keys) == 0 {
		return nil, nil
	}
	if err := this.Client.Del(keys...).Err(); err != nil {
		return nil, err
	}
	return nil, nil
}
