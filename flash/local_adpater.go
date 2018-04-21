package flash

import "github.com/sanksons/flashcache/flash/common"

type Local struct {
}

func (this *Local) Set(i common.CacheItem) error {
	return nil
}

func (this *Local) Get(key string) ([]byte, error) {
	return nil, nil
}

func (this *Local) MGet(a ...string) (map[string][]byte, error) {
	return nil, nil
}
