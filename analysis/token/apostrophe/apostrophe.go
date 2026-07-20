package apostrophe

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "apostrophe"

const RightSingleQuotationMark = "’"
const Apostrophe = "'"
const Apostrophes = Apostrophe + RightSingleQuotationMark

type ApostropheFilter struct{}

func NewApostropheFilter() *ApostropheFilter { _ = "STUB: not implemented"; return nil }

func (s *ApostropheFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func ApostropheFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, ApostropheFilterConstructor)
	if err != nil {
		panic(err)
	}
}
