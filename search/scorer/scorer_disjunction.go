package scorer

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
)

var reflectStaticSizeDisjunctionQueryScorer int

func init() {
	var dqs DisjunctionQueryScorer
	reflectStaticSizeDisjunctionQueryScorer = int(reflect.TypeOf(dqs).Size())
}

type DisjunctionQueryScorer struct {
	options search.SearcherOptions
}

func (s *DisjunctionQueryScorer) Size() int { _ = "STUB: not implemented"; return 0 }

func NewDisjunctionQueryScorer(options search.SearcherOptions) *DisjunctionQueryScorer {
	_ = "STUB: not implemented"
	return nil
}

func (s *DisjunctionQueryScorer) Score(ctx *search.SearchContext, constituents []*search.DocumentMatch, countMatch, countTotal int) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *DisjunctionQueryScorer) ScoreAndExplBreakdown(ctx *search.SearchContext, constituents []*search.DocumentMatch,
	matchingIdxs []int, originalPositions []int, countTotal int) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}
