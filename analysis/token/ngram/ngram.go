package ngram

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "ngram"

type NgramFilter struct {
	minLength int
	maxLength int
}

func NewNgramFilter(minLength, maxLength int) *NgramFilter { _ = "STUB: not implemented"; return nil }

func (s *NgramFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func NgramFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, NgramFilterConstructor)
	if err != nil {
		panic(err)
	}
}

func convertToInt(val interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }
