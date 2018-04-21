package main

import (
	"fmt"
	"log"

	"github.com/sanksons/flashcache/flash/common"
	fredis "github.com/sanksons/flashcache/flash/redis"

	"github.com/sanksons/flashcache/flash"

	"github.com/go-redis/redis"
)

const CACHING_ENGINE = "redis-simple"

func main() {
	fmt.Println("I am main")
	cacheAdapter := Initialize()
	err := cacheAdapter.Set(common.CacheItem{Key: "A", Value: []byte("hello1")})
	if err != nil {
		log.Fatal(err)
	}
	val, err := cacheAdapter.MGet([]string{"C", "A", "B", "q"}...)
	if err != nil && err == common.KeyNotExists {
		log.Fatal("key does not exists")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", val)

}

func Initialize() flash.CacheAdapter {
	switch CACHING_ENGINE {
	case flash.ADAPTER_TYPE_REDIS_SIMPLE:
		rs := new(fredis.RedisSimple)
		rs.Initialize(fredis.RedisSimpleConfig{Addr: "localhost:6379"})
		return rs
	case flash.ADAPTER_TYPE_REDIS_CLUSTER:
		return nil
	case flash.ADAPTER_TYPE_LOCAL:
		return nil
	}
	return nil
}

// func InitializeRedis() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	rs := flash.RedisSimple{
// 		client,
// 	}
// 	r := flash.Redis{
// 		Client: &rs,
// 	}
// 	pong, err := client.Ping().Result()
// 	fmt.Println(pong, err)
// 	r.Set(flash.CacheItem{})
// }

func InitializeRedisCluster() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	client.Ping()

	pong, err := client.Ping().Result()

	fmt.Println(pong, err)

}
