package fr

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_fr_snowball"

type FrenchStemmerFilter struct {
}

func NewFrenchStemmerFilter() *FrenchStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *FrenchStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func FrenchStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, FrenchStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
