package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/joaojeronimo/go-crc16"
	"github.com/sanksons/tavern/utils"
)

const SLOTS_COUNT = 16384

type RedisClusterConfig struct {
	Addrs []string
}

type RedisCluster struct {
	redisbase
}

func InitializeRedisCluster(config RedisClusterConfig) *RedisCluster {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: config.Addrs,
	})
	redisCluster := new(RedisCluster)
	redisCluster.Client = client
	return redisCluster
}

func (this *RedisCluster) slotify(keys ...utils.CacheKey) map[uint16][]utils.CacheKey {
	m := make(map[uint16][]utils.CacheKey)
	for _, key := range keys {
		slot := (crc16.Crc16([]byte(key))) % SLOTS_COUNT
		if _, ok := m[slot]; !ok {
			m[slot] = []utils.CacheKey{key}
		} else {
			m[slot] = append(m[slot], key)
		}
	}
	return m
}

// Override base redis implementation of Mget.
// @done: - Group keys based on the slots they belong.
// @done: - Make concurrent calls to attain parallelism.
func (this *RedisCluster) MGet(a ...utils.CacheKey) (map[utils.CacheKey][]byte, error) {
	if len(a) == 0 {
		return nil, nil
	}
	//slotify
	m := this.slotify(a...)

	//parallelize
	ch := make(chan map[utils.CacheKey][]byte)
	max := len(m)
	for _, keys := range m { //parrallelize calls to redis
		go func(keys ...utils.CacheKey) {
			keydata, _ := this.redisbase.MGet(keys...)
			ch <- keydata
		}(keys...)
	}

	datamap := make(map[utils.CacheKey][]byte)
	// join all
	var count int
	for {
		if count == max {
			break
		}
		count++
		chunkdata := <-ch
		for k, v := range chunkdata {
			datamap[k] = v
		}
	}
	return datamap, nil
}

//
// Override base redis implementation of Mset
//
func (this *RedisCluster) MSet(items ...utils.CacheItem) (map[utils.CacheKey]bool, error) {
	if len(items) == 0 {
		return nil, nil
	}
	keys := make([]utils.CacheKey, 0, len(items))
	mitems := make(map[utils.CacheKey]utils.CacheItem)
	for _, item := range items {
		keys = append(keys, item.Key)
		mitems[item.Key] = item
	}

	//slotify
	m := this.slotify(keys...)

	//parallelize
	funcs := make([]func() interface{}, 0, len(m))
	for _, keys := range m { //parrallelize calls to redis
		cacheItems := make([]utils.CacheItem, 0, len(keys))
		for _, key := range keys {
			cacheItems = append(cacheItems, mitems[key])
		}
		//prepare functions to be executed parallely.
		funcs = append(funcs, func(cacheItems ...utils.CacheItem) func() interface{} {
			return func() interface{} {
				keydata, _ := this.redisbase.MSet(cacheItems...) //this is local keys.
				return keydata
			}
		}(cacheItems...))

	}
	datamap := make(map[utils.CacheKey]bool)
	results := utils.Parallelize(funcs)

	//club all
	for _, i := range results {
		datamapTmp, ok := i.(map[utils.CacheKey]bool)
		if !ok {
			fmt.Printf("Expected map[utils.CacheKey]bool, found: %T", i)
		}
		for k, v := range datamapTmp {
			datamap[k] = v
		}
	}
	return datamap, nil
}

//
// Override base redis implementation of Destroy
//
func (this *RedisCluster) Destroy(keys ...utils.CacheKey) (map[utils.CacheKey]bool, error) {

	if len(keys) == 0 {
		return nil, nil
	}
	//slotify
	m := this.slotify(keys...)

	//parallelize
	funcs := make([]func() interface{}, 0, len(m))
	for _, keys := range m { //parrallelize calls to redis

		funcs = append(funcs, func(keys ...utils.CacheKey) func() interface{} {
			return func() interface{} {
				keydata, _ := this.redisbase.Destroy(keys...) //this is local keys.
				return keydata
			}
		}(keys...))

	}
	datamap := make(map[utils.CacheKey]bool)
	results := utils.Parallelize(funcs)

	//club all
	for _, i := range results {
		datamapTmp, ok := i.(map[utils.CacheKey]bool)
		if !ok {
			fmt.Printf("Expected map[string]bool, found: %T", i)
		}
		for k, v := range datamapTmp {
			datamap[k] = v
		}
	}
	return datamap, nil
}
