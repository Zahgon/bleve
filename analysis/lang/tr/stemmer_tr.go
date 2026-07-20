package tr

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_tr_snowball"

type TurkishStemmerFilter struct {
}

func NewTurkishStemmerFilter() *TurkishStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *TurkishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func TurkishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, TurkishStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
