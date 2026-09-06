package registry

import (
	"github.com/blevesearch/bleve/v2/search/highlight"
)

func RegisterHighlighter(name string, constructor HighlighterConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type HighlighterConstructor func(config map[string]interface{}, cache *Cache) (highlight.Highlighter, error)
type HighlighterRegistry map[string]HighlighterConstructor

type HighlighterCache struct {
	*ConcurrentCache
}

func NewHighlighterCache() *HighlighterCache { _ = "STUB: not implemented"; return nil }

func HighlighterBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *HighlighterCache) HighlighterNamed(name string, cache *Cache) (highlight.Highlighter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Highlighter), nil
}

func (c *HighlighterCache) DefineHighlighter(name string, typ string, config map[string]interface{}, cache *Cache) (highlight.Highlighter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Highlighter), nil
}

func HighlighterTypesAndInstances() ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
