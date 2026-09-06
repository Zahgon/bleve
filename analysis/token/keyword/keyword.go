package keyword

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "keyword_marker"

type KeyWordMarkerFilter struct {
	keyWords analysis.TokenMap
}

func NewKeyWordMarkerFilter(keyWords analysis.TokenMap) *KeyWordMarkerFilter {
	_ = "STUB: not implemented"
	return nil
}

func (f *KeyWordMarkerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func KeyWordMarkerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, KeyWordMarkerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
