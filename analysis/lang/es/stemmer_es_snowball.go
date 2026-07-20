package es

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_es_snowball"

type SpanishStemmerFilter struct {
}

func NewSpanishStemmerFilter() *SpanishStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *SpanishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func SpanishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, SpanishStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
