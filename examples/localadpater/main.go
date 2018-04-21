package main

import (
	"fmt"
	"log"

	"github.com/sanksons/tavern/adapters/local"
	"github.com/sanksons/tavern/utils"
)

const CACHING_ENGINE = "redis-simple"

func main() {
	cacheAdapter := local.Initialize(local.LocalAdapterConfig{})
	cacheAdapter.Set(utils.CacheItem{Key: "A", Value: []byte("I am A")})

	data, err := cacheAdapter.Get("A")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n", data)
}
