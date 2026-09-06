package length

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "length"

type LengthFilter struct {
	min int
	max int
}

func NewLengthFilter(min, max int) *LengthFilter { _ = "STUB: not implemented"; return nil }

func (f *LengthFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func LengthFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, LengthFilterConstructor)
	if err != nil {
		panic(err)
	}
}
