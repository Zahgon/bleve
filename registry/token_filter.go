package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func RegisterTokenFilter(name string, constructor TokenFilterConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type TokenFilterConstructor func(config map[string]interface{}, cache *Cache) (analysis.TokenFilter, error)
type TokenFilterRegistry map[string]TokenFilterConstructor

type TokenFilterCache struct {
	*ConcurrentCache
}

func NewTokenFilterCache() *TokenFilterCache { _ = "STUB: not implemented"; return nil }

func TokenFilterBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TokenFilterCache) TokenFilterNamed(name string, cache *Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func (c *TokenFilterCache) DefineTokenFilter(name string, typ string, config map[string]interface{}, cache *Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func TokenFilterTypesAndInstances() ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
