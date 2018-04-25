package redis

import (
	"github.com/go-redis/redis"
	"github.com/joaojeronimo/go-crc16"
)

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

// Override base redis implementation of Mget.
// @done: - Group keys based on the slots they belong.
// @todo: - Make concurrent calls to attain parallelism.
func (this *RedisCluster) MGet(a ...string) (map[string][]byte, error) {

	if len(a) == 0 {
		return nil, nil
	}
	//slotify
	m := make(map[uint16][]string)
	for _, key := range a {
		slot := (crc16.Crc16([]byte(key))) % SLOTS_COUNT
		if _, ok := m[slot]; !ok {
			m[slot] = []string{key}
		} else {
			m[slot] = append(m[slot], key)
		}
	}

	ch := make(chan map[string][]byte)
	max := len(m)
	for _, keys := range m { //parrallelize calls to redis
		go func(keys ...string) {
			keydata, _ := this.redisbase.MGet(keys...)
			ch <- keydata
		}(keys...)
	}

	datamap := make(map[string][]byte)
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
