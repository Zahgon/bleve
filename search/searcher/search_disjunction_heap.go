package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeDisjunctionHeapSearcher int
var reflectStaticSizeSearcherCurr int

func init() {
	var dhs DisjunctionHeapSearcher
	reflectStaticSizeDisjunctionHeapSearcher = int(reflect.TypeOf(dhs).Size())

	var sc SearcherCurr
	reflectStaticSizeSearcherCurr = int(reflect.TypeOf(sc).Size())
}

type SearcherCurr struct {
	searcher    search.Searcher
	curr        *search.DocumentMatch
	matchingIdx int
}

type DisjunctionHeapSearcher struct {
	indexReader index.IndexReader

	numSearchers           int
	scorer                 *scorer.DisjunctionQueryScorer
	min                    int
	queryNorm              float64
	retrieveScoreBreakdown bool
	initialized            bool
	searchers              []search.Searcher
	heap                   []*SearcherCurr

	matching      []*search.DocumentMatch
	matchingIdxs  []int
	matchingCurrs []*SearcherCurr

	bytesRead uint64
}

func newDisjunctionHeapSearcher(ctx context.Context, indexReader index.IndexReader,
	searchers []search.Searcher, min float64, options search.SearcherOptions,
	limit bool) (
	*DisjunctionHeapSearcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DisjunctionHeapSearcher) computeQueryNorm() { _ = "STUB: not implemented"; return }

func (s *DisjunctionHeapSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionHeapSearcher) initSearchers(ctx *search.SearchContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *DisjunctionHeapSearcher) updateMatches() error { _ = "STUB: not implemented"; return nil }

func (s *DisjunctionHeapSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionHeapSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *DisjunctionHeapSearcher) Next(ctx *search.SearchContext) (
	*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DisjunctionHeapSearcher) Advance(ctx *search.SearchContext,
	ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DisjunctionHeapSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionHeapSearcher) Close() (rv error) { _ = "STUB: not implemented"; return nil }

func (s *DisjunctionHeapSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionHeapSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionHeapSearcher) Optimize(kind string, octx index.OptimizableContext) (
	index.OptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.OptimizableContext), nil
}

func (s *DisjunctionHeapSearcher) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *DisjunctionHeapSearcher) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s *DisjunctionHeapSearcher) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s *DisjunctionHeapSearcher) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (s *DisjunctionHeapSearcher) Pop() interface{} { _ = "STUB: not implemented"; return nil }
