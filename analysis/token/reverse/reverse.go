package reverse

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "reverse"

type ReverseFilter struct {
}

func NewReverseFilter() *ReverseFilter { _ = "STUB: not implemented"; return nil }

func (f *ReverseFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func ReverseFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, ReverseFilterConstructor)
	if err != nil {
		panic(err)
	}
}

func reverse(s []byte) []byte { _ = "STUB: not implemented"; return nil }
