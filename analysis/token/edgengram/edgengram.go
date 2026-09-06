package edgengram

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "edge_ngram"

type Side bool

const BACK Side = true
const FRONT Side = false

type EdgeNgramFilter struct {
	back      Side
	minLength int
	maxLength int
}

func NewEdgeNgramFilter(side Side, minLength, maxLength int) *EdgeNgramFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *EdgeNgramFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func EdgeNgramFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, EdgeNgramFilterConstructor)
	if err != nil {
		panic(err)
	}
}
