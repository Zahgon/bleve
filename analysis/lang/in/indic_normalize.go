package in

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const NormalizeName = "normalize_in"

type IndicNormalizeFilter struct {
}

func NewIndicNormalizeFilter() *IndicNormalizeFilter { _ = "STUB: not implemented"; return nil }

func (s *IndicNormalizeFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func NormalizerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(NormalizeName, NormalizerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
