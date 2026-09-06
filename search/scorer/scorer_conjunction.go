package scorer

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
)

var reflectStaticSizeConjunctionQueryScorer int

func init() {
	var cqs ConjunctionQueryScorer
	reflectStaticSizeConjunctionQueryScorer = int(reflect.TypeOf(cqs).Size())
}

type ConjunctionQueryScorer struct {
	options search.SearcherOptions
}

func (s *ConjunctionQueryScorer) Size() int { _ = "STUB: not implemented"; return 0 }

func NewConjunctionQueryScorer(options search.SearcherOptions) *ConjunctionQueryScorer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ConjunctionQueryScorer) Score(ctx *search.SearchContext, constituents []*search.DocumentMatch) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}
