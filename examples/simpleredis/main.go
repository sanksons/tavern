package main

import (
	"fmt"
	"log"

	"github.com/sanksons/tavern/utils"

	"github.com/sanksons/tavern/adapters/redis"
)

const CACHING_ENGINE = "redis-simple"

func main() {

	//This is how we initialize redis adapter
	cacheAdapter := redis.InitializeRedisSimple(redis.RedisSimpleConfig{
		Addr: "localhost:6379",
	})

	//This is how we set a key
	cacheAdapter.Set(utils.CacheItem{Key: "A", Value: []byte("I am A")})

	//This is how we get a key
	data, err := cacheAdapter.Get("A")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n", data)

	//This is how we set multiple keys
	items := prepareCacheItems()
	fmt.Println("\nSet multiple items:")
	result, err := cacheAdapter.MSet(items...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \n%+v\n", result)

	//This is how we get multiple keys
	fmt.Println("\n get multiple Items:")
	resultget, err := cacheAdapter.MGet("A", "B", "C", "D")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: \n%+v\n", resultget)

	//this is how we destry keys
	fmt.Println("\n delete Items:")
	resultdelete, err := cacheAdapter.Destroy("A", "B", "C")
	fmt.Printf("Result: \n%+v\n", resultdelete)

}

func prepareCacheItems() []utils.CacheItem {
	data := map[string]string{
		"A": "I am A",
		"B": "I am A",
		"C": "I am C",
	}
	cacheItems := make([]utils.CacheItem, 0)
	for k, v := range data {
		item := utils.CacheItem{
			Key:   utils.CacheKey(k),
			Value: []byte(v),
		}
		cacheItems = append(cacheItems, item)
	}
	return cacheItems
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
