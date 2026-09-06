package it

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_it_snowball"

type ItalianStemmerFilter struct {
}

func NewItalianStemmerFilter() *ItalianStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *ItalianStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func ItalianStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, ItalianStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
