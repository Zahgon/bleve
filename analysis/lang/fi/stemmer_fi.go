package fi

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_fi_snowball"

type FinnishStemmerFilter struct {
}

func NewFinnishStemmerFilter() *FinnishStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *FinnishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func FinnishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, FinnishStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
