package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func RegisterTokenMap(name string, constructor TokenMapConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type TokenMapConstructor func(config map[string]interface{}, cache *Cache) (analysis.TokenMap, error)
type TokenMapRegistry map[string]TokenMapConstructor

type TokenMapCache struct {
	*ConcurrentCache
}

func NewTokenMapCache() *TokenMapCache { _ = "STUB: not implemented"; return nil }

func TokenMapBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TokenMapCache) TokenMapNamed(name string, cache *Cache) (analysis.TokenMap, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenMap), nil
}

func (c *TokenMapCache) DefineTokenMap(name string, typ string, config map[string]interface{}, cache *Cache) (analysis.TokenMap, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenMap), nil
}

func TokenMapTypesAndInstances() ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }
