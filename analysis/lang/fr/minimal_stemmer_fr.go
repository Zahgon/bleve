package fr

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const MinimalStemmerName = "stemmer_fr_min"

type FrenchMinimalStemmerFilter struct {
}

func NewFrenchMinimalStemmerFilter() *FrenchMinimalStemmerFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *FrenchMinimalStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func minstem(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func FrenchMinimalStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(MinimalStemmerName, FrenchMinimalStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
