package no

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_no_snowball"

type NorwegianStemmerFilter struct {
}

func NewNorwegianStemmerFilter() *NorwegianStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *NorwegianStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func NorwegianStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, NorwegianStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
