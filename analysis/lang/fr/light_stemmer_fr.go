package fr

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const LightStemmerName = "stemmer_fr_light"

type FrenchLightStemmerFilter struct {
}

func NewFrenchLightStemmerFilter() *FrenchLightStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *FrenchLightStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func stem(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func norm(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func FrenchLightStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(LightStemmerName, FrenchLightStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
