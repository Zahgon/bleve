package unique

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "unique"

type UniqueTermFilter struct{}

func NewUniqueTermFilter() *UniqueTermFilter { _ = "STUB: not implemented"; return nil }

func (f *UniqueTermFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func UniqueTermFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, UniqueTermFilterConstructor)
	if err != nil {
		panic(err)
	}
}
