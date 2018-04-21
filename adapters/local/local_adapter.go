package local

import (
	"github.com/sanksons/tavern/utils"
)

type Local struct {
}

func (this *Local) Set(i utils.CacheItem) error {
	return nil
}

func (this *Local) Get(key string) ([]byte, error) {
	return nil, nil
}

func (this *Local) MGet(a ...string) (map[string][]byte, error) {
	return nil, nil
}
