package lowercase

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "to_lower"

type LowerCaseFilter struct {
}

func NewLowerCaseFilter() *LowerCaseFilter { _ = "STUB: not implemented"; return nil }

func (f *LowerCaseFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func LowerCaseFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, LowerCaseFilterConstructor)
	if err != nil {
		panic(err)
	}
}

func toLowerDeferredCopy(s []byte) []byte { _ = "STUB: not implemented"; return nil }
