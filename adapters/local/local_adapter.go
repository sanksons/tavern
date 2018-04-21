package local

import (
	"fmt"

	localcache "github.com/patrickmn/go-cache"
	"github.com/sanksons/tavern/utils"
)

type LocalAdapterConfig struct {
}

type Local struct {
	Client *localcache.Cache
}

func (this *Local) Set(i utils.CacheItem) error {
	this.Client.Set(string(i.Key), i.Value, i.Expiration)
	return nil
}

func (this *Local) Get(key string) ([]byte, error) {
	data, found := this.Client.Get(key)
	if !found {
		return nil, utils.KeyNotExists
	}
	databytes, ok := data.([]byte)
	if !ok {
		return nil, fmt.Errorf("malformed key data")
	}
	return databytes, nil
}

func (this *Local) MGet(a ...string) (map[string][]byte, error) {
	return nil, nil
}
