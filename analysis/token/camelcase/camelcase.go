package camelcase

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "camelCase"

type CamelCaseFilter struct{}

func NewCamelCaseFilter() *CamelCaseFilter { _ = "STUB: not implemented"; return nil }

func (f *CamelCaseFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func CamelCaseFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, CamelCaseFilterConstructor)
	if err != nil {
		panic(err)
	}
}
