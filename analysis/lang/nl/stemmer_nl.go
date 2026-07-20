package nl

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_nl_snowball"

type DutchStemmerFilter struct {
}

func NewDutchStemmerFilter() *DutchStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *DutchStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func DutchStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, DutchStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
