package registry

import (
	"github.com/blevesearch/bleve/v2/search/highlight"
)

func RegisterFragmentFormatter(name string, constructor FragmentFormatterConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type FragmentFormatterConstructor func(config map[string]interface{}, cache *Cache) (highlight.FragmentFormatter, error)
type FragmentFormatterRegistry map[string]FragmentFormatterConstructor

type FragmentFormatterCache struct {
	*ConcurrentCache
}

func NewFragmentFormatterCache() *FragmentFormatterCache { _ = "STUB: not implemented"; return nil }

func FragmentFormatterBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FragmentFormatterCache) FragmentFormatterNamed(name string, cache *Cache) (highlight.FragmentFormatter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.FragmentFormatter), nil
}

func (c *FragmentFormatterCache) DefineFragmentFormatter(name string, typ string, config map[string]interface{}, cache *Cache) (highlight.FragmentFormatter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.FragmentFormatter), nil
}

func FragmentFormatterTypesAndInstances() ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
