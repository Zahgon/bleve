package da

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_da_snowball"

type DanishStemmerFilter struct {
}

func NewDanishStemmerFilter() *DanishStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *DanishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func DanishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, DanishStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
