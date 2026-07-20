package pl

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/analysis/lang/pl/stempel"
	"github.com/blevesearch/bleve/v2/registry"
)

const SnowballStemmerName = "stemmer_pl"

type PolishStemmerFilter struct {
	trie stempel.Trie
}

func NewPolishStemmerFilter() (*PolishStemmerFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PolishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func PolishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(SnowballStemmerName, PolishStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
