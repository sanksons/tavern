package utils

import "time"

type CacheKey string

type CacheItem struct {
	Key        CacheKey
	Value      []byte
	Expiration time.Duration
}

type CacheItems []*CacheItem
