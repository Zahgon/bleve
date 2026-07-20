package es

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const LightStemmerName = "stemmer_es_light"

type SpanishLightStemmerFilter struct {
}

func NewSpanishLightStemmerFilter() *SpanishLightStemmerFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpanishLightStemmerFilter) Filter(
	input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func stem(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func SpanishLightStemmerFilterConstructor(config map[string]interface{},
	cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(LightStemmerName, SpanishLightStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
