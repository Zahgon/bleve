package html

import (
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search/highlight"
)

const Name = "html"

func Constructor(config map[string]interface{}, cache *registry.Cache) (highlight.Highlighter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Highlighter), nil
}

func init() {
	err := registry.RegisterHighlighter(Name, Constructor)
	if err != nil {
		panic(err)
	}
}
