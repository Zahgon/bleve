package porter

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "stemmer_porter"

type PorterStemmer struct {
}

func NewPorterStemmer() *PorterStemmer { _ = "STUB: not implemented"; return nil }

func (s *PorterStemmer) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func PorterStemmerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, PorterStemmerConstructor)
	if err != nil {
		panic(err)
	}
}
