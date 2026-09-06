package snowball

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "stemmer_snowball"

type SnowballStemmer struct {
	language string
}

func NewSnowballStemmer(language string) *SnowballStemmer { _ = "STUB: not implemented"; return nil }

func (s *SnowballStemmer) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func SnowballStemmerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, SnowballStemmerConstructor)
	if err != nil {
		panic(err)
	}
}
