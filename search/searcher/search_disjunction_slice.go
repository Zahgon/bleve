package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeDisjunctionSliceSearcher int

func init() {
	var ds DisjunctionSliceSearcher
	reflectStaticSizeDisjunctionSliceSearcher = int(reflect.TypeOf(ds).Size())
}

type DisjunctionSliceSearcher struct {
	indexReader            index.IndexReader
	searchers              []search.Searcher
	originalPos            []int
	numSearchers           int
	queryNorm              float64
	retrieveScoreBreakdown bool
	currs                  []*search.DocumentMatch
	scorer                 *scorer.DisjunctionQueryScorer
	min                    int
	matching               []*search.DocumentMatch
	matchingIdxs           []int
	initialized            bool
	bytesRead              uint64
}

func newDisjunctionSliceSearcher(ctx context.Context, indexReader index.IndexReader,
	qsearchers []search.Searcher, min float64, options search.SearcherOptions,
	limit bool) (
	*DisjunctionSliceSearcher, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DisjunctionSliceSearcher) computeQueryNorm() { _ = "STUB: not implemented"; return }

func (s *DisjunctionSliceSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionSliceSearcher) initSearchers(ctx *search.SearchContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DisjunctionSliceSearcher) updateMatches() error { _ = "STUB: not implemented"; return nil }

func (s *DisjunctionSliceSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionSliceSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *DisjunctionSliceSearcher) Next(ctx *search.SearchContext) (
	*search.DocumentMatch, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DisjunctionSliceSearcher) Advance(ctx *search.SearchContext,
	ID index.IndexInternalID,
) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DisjunctionSliceSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionSliceSearcher) Close() (rv error) { _ = "STUB: not implemented"; return nil }

func (s *DisjunctionSliceSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionSliceSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionSliceSearcher) Optimize(kind string, octx index.OptimizableContext) (
	index.OptimizableContext, error,
) {
	_ = "STUB: not implemented"
	return *new(index.OptimizableContext), nil
}
