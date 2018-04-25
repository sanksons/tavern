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

func (this *RedisCluster) Slotify(keys ...string) map[uint16][]string {
	m := make(map[uint16][]string)
	for _, key := range keys {
		slot := (crc16.Crc16([]byte(key))) % SLOTS_COUNT
		if _, ok := m[slot]; !ok {
			m[slot] = []string{key}
		} else {
			m[slot] = append(m[slot], key)
		}
	}
	return m
}

// Override base redis implementation of Mget.
// @done: - Group keys based on the slots they belong.
// @done: - Make concurrent calls to attain parallelism.
func (this *RedisCluster) MGet(a ...string) (map[string][]byte, error) {

	if len(a) == 0 {
		return nil, nil
	}
	//slotify
	m := this.Slotify(a...)

	//parallelize

	ch := make(chan map[string][]byte)
	max := len(m)
	for _, keys := range m { //parrallelize calls to redis
		go func(keys ...string) {
			keydata, _ := this.redisbase.MGet(keys...)
			ch <- keydata
		}(keys...)
	}

	datamap := make(map[string][]byte)
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

func (this *RedisCluster) MSet(items ...utils.CacheItem) (map[string]bool, error) {
	return nil, nil
}

// Override base redis implementation of Destroy

func (this *RedisCluster) Destroy(keys ...string) (map[string]bool, error) {

	if len(keys) == 0 {
		return nil, nil
	}
	//slotify
	m := this.Slotify(keys...)

	//parallelize
	funcs := make([]func() interface{}, 0)
	for _, keys := range m { //parrallelize calls to redis

		funcs = append(funcs, func(keys ...string) func() interface{} {
			return func() interface{} {
				keydata, _ := this.redisbase.Destroy(keys...) //this is local keys.
				return keydata
			}
		}(keys...))

	}
	datamap := make(map[string]bool)
	results := utils.Parallelize(funcs)
	for _, i := range results {
		datamapTmp, ok := i.(map[string]bool)
		if !ok {
			fmt.Printf("Expected map[string]bool, found: %T", i)
		}
		for k, v := range datamapTmp {
			datamap[k] = v
		}
	}
	return datamap, nil
}
