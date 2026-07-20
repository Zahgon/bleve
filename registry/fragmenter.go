package registry

import (
	"github.com/blevesearch/bleve/v2/search/highlight"
)

func RegisterFragmenter(name string, constructor FragmenterConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type FragmenterConstructor func(config map[string]interface{}, cache *Cache) (highlight.Fragmenter, error)
type FragmenterRegistry map[string]FragmenterConstructor

type FragmenterCache struct {
	*ConcurrentCache
}

func NewFragmenterCache() *FragmenterCache { _ = "STUB: not implemented"; return nil }

func FragmenterBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *FragmenterCache) FragmenterNamed(name string, cache *Cache) (highlight.Fragmenter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Fragmenter), nil
}

func (c *FragmenterCache) DefineFragmenter(name string, typ string, config map[string]interface{}, cache *Cache) (highlight.Fragmenter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Fragmenter), nil
}

func FragmenterTypesAndInstances() ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }
