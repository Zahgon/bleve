package de

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_de_snowball"

type GermanStemmerFilter struct {
}

func NewGermanStemmerFilter() *GermanStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *GermanStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func GermanStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, GermanStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
