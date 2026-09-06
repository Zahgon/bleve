package ru

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_ru_snowball"

type RussianStemmerFilter struct {
}

func NewRussianStemmerFilter() *RussianStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *RussianStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func RussianStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, RussianStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
