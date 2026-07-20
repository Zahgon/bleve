package cjk

import (
	"container/ring"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const BigramName = "cjk_bigram"

type CJKBigramFilter struct {
	outputUnigram bool
}

func NewCJKBigramFilter(outputUnigram bool) *CJKBigramFilter { _ = "STUB: not implemented"; return nil }

func (s *CJKBigramFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func (s *CJKBigramFilter) flush(r *ring.Ring, itemsInRing *int, pos int) *analysis.Token {
	_ = "STUB: not implemented"
	return nil
}

func (s *CJKBigramFilter) outputBigram(r *ring.Ring, itemsInRing *int, pos int) *analysis.Token {
	_ = "STUB: not implemented"
	return nil
}

func (s *CJKBigramFilter) buildUnigram(r *ring.Ring, itemsInRing *int, pos int) *analysis.Token {
	_ = "STUB: not implemented"
	return nil
}

func CJKBigramFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(BigramName, CJKBigramFilterConstructor)
	if err != nil {
		panic(err)
	}
}
