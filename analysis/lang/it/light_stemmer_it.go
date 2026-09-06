package it

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const LightStemmerName = "stemmer_it_light"

type ItalianLightStemmerFilter struct {
}

func NewItalianLightStemmerFilterFilter() *ItalianLightStemmerFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ItalianLightStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func stem(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func ItalianLightStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(LightStemmerName, ItalianLightStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
