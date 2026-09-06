package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func RegisterCharFilter(name string, constructor CharFilterConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type CharFilterConstructor func(config map[string]interface{}, cache *Cache) (analysis.CharFilter, error)
type CharFilterRegistry map[string]CharFilterConstructor

type CharFilterCache struct {
	*ConcurrentCache
}

func NewCharFilterCache() *CharFilterCache { _ = "STUB: not implemented"; return nil }

func CharFilterBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CharFilterCache) CharFilterNamed(name string, cache *Cache) (analysis.CharFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.CharFilter), nil
}

func (c *CharFilterCache) DefineCharFilter(name string, typ string, config map[string]interface{}, cache *Cache) (analysis.CharFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.CharFilter), nil
}

func CharFilterTypesAndInstances() ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }
