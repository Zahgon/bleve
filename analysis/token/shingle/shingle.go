package shingle

import (
	"container/ring"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "shingle"

type ShingleFilter struct {
	min            int
	max            int
	outputOriginal bool
	tokenSeparator string
	fill           string
}

func NewShingleFilter(min, max int, outputOriginal bool, sep, fill string) *ShingleFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ShingleFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func (s *ShingleFilter) shingleCurrentRingState(ring *ring.Ring, itemsInRing int) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func ShingleFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, ShingleFilterConstructor)
	if err != nil {
		panic(err)
	}
}
