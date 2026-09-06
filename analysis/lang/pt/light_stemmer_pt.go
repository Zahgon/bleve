package pt

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const LightStemmerName = "stemmer_pt_light"

type PortugueseLightStemmerFilter struct {
}

func NewPortugueseLightStemmerFilter() *PortugueseLightStemmerFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PortugueseLightStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func stem(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func removeSuffix(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func normFeminine(input []rune) []rune { _ = "STUB: not implemented"; return nil }

func PortugueseLightStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(LightStemmerName, PortugueseLightStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
