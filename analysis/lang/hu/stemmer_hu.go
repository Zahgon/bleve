package hu

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_hu_snowball"

type HungarianStemmerFilter struct {
}

func NewHungarianStemmerFilter() *HungarianStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *HungarianStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func HungarianStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, HungarianStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
