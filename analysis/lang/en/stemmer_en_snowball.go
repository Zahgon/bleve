package en

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_en_snowball"

type EnglishStemmerFilter struct {
}

func NewEnglishStemmerFilter() *EnglishStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *EnglishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func EnglishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, EnglishStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
