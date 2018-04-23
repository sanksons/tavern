package main

import (
	"github.com/sanksons/tavern/utils"

	"github.com/sanksons/tavern/adapters/redis"

	"github.com/sanksons/tavern"
)

const CACHING_ENGINE = "redis-simple"

func main() {
	//	client := goredis.NewClient(&goredis.Options{
	//		Addr:     "localhost:6379",
	//		Password: "", // no password set
	//		DB:       0,  // use default DB
	//	})

	// err := client.MSet([]string{"key1", "value1", "key2", "value2", "key3", "value3"}).Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 	fmt.Println("I am main")
	cacheAdapter := Initialize()
	cacheAdapter.MSet([]utils.CacheItem{
		utils.CacheItem{
			Key:   "key10",
			Value: []byte("I am key 10"),
		},
		utils.CacheItem{
			Key:   "key11",
			Value: []byte("I am key 11"),
		},
	}...)

	cacheAdapter.Destroy("key10", "key11")

}

func Initialize() tavern.CacheAdapter {
	switch CACHING_ENGINE {
	case tavern.ADAPTER_TYPE_REDIS_SIMPLE:

		rs := new(redis.RedisSimple)
		rs.Initialize(redis.RedisSimpleConfig{Addr: "localhost:6379"})
		return rs
	case tavern.ADAPTER_TYPE_REDIS_CLUSTER:
		return nil
	case tavern.ADAPTER_TYPE_LOCAL:
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

// func InitializeRedisCluster() {
// 	client := redis.NewClusterClient(&redis.ClusterOptions{
// 		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
// 	})
// 	client.Ping()

// 	pong, err := client.Ping().Result()

// 	fmt.Println(pong, err)

// }
