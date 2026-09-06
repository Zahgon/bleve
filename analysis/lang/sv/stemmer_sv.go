package sv

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_sv_snowball"

type SwedishStemmerFilter struct {
}

func NewSwedishStemmerFilter() *SwedishStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *SwedishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func SwedishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, SwedishStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
