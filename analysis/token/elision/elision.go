package elision

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "elision"

const RightSingleQuotationMark = '’'
const Apostrophe = '\''

type ElisionFilter struct {
	articles analysis.TokenMap
}

func NewElisionFilter(articles analysis.TokenMap) *ElisionFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ElisionFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func ElisionFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, ElisionFilterConstructor)
	if err != nil {
		panic(err)
	}
}
