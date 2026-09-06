package truncate

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "truncate_token"

type TruncateTokenFilter struct {
	length int
}

func NewTruncateTokenFilter(length int) *TruncateTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *TruncateTokenFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func TruncateTokenFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, TruncateTokenFilterConstructor)
	if err != nil {
		panic(err)
	}
}
