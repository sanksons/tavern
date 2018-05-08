package entity

import "time"

//this is the key as supplied by user
type CacheKey string

func (this CacheKey) GetMachineKey() MachineKey {
	return MachineKey(this)
}

func (this CacheKey) String() string {
	return string(this)
}

//this is the name of key as stored in cache system.
type MachineKey string

func (this MachineKey) GetCacheKey() CacheKey {
	return CacheKey(this)
}

func (this MachineKey) String() string {
	return string(this)
}

type CacheItem struct {
	Key        CacheKey
	Value      []byte
	Expiration time.Duration //in seconds
}
