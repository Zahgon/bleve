package ro

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_ro_snowball"

type RomanianStemmerFilter struct {
}

func NewRomanianStemmerFilter() *RomanianStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *RomanianStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func RomanianStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, RomanianStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
