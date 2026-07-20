package registry

import (
	"fmt"
	"sync"
)

var ErrAlreadyDefined = fmt.Errorf("item already defined")

type CacheBuild func(name string, config map[string]interface{}, cache *Cache) (interface{}, error)

type ConcurrentCache struct {
	mutex sync.RWMutex
	data  map[string]interface{}
}

func NewConcurrentCache() *ConcurrentCache { _ = "STUB: not implemented"; return nil }

func (c *ConcurrentCache) ItemNamed(name string, cache *Cache, build CacheBuild) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConcurrentCache) DefineItem(name string, typ string, config map[string]interface{}, cache *Cache, build CacheBuild) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
