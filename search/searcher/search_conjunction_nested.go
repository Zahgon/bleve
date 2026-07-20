package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeNestedConjunctionSearcher int

func init() {
	var ncs NestedConjunctionSearcher
	reflectStaticSizeNestedConjunctionSearcher = int(reflect.TypeOf(ncs).Size())
}

type NestedConjunctionSearcher struct {
	nestedReader  index.NestedReader
	searchers     []search.Searcher
	queryNorm     float64
	currs         []*search.DocumentMatch
	currAncestors [][]index.AncestorID
	currKeys      []index.AncestorID
	initialized   bool
	joinIdx       int
	options       search.SearcherOptions
	docQueue      *CoalesceQueue

	advanceID index.IndexInternalID

	ancestors []index.AncestorID
}

func NewNestedConjunctionSearcher(ctx context.Context, indexReader index.IndexReader,
	searchers []search.Searcher, joinIdx int, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (s *NestedConjunctionSearcher) computeQueryNorm() { _ = "STUB: not implemented"; return }

func (s *NestedConjunctionSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *NestedConjunctionSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *NestedConjunctionSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *NestedConjunctionSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *NestedConjunctionSearcher) Close() (rv error) { _ = "STUB: not implemented"; return nil }

func (s *NestedConjunctionSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *NestedConjunctionSearcher) DocumentMatchPoolSize() int {
	_ = "STUB: not implemented"
	return 0
}

func (s *NestedConjunctionSearcher) initialize(ctx *search.SearchContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *NestedConjunctionSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ancestorFromRoot(ancestors []index.AncestorID, pos int) index.AncestorID {
	_ = "STUB: not implemented"
	return *new(index.AncestorID)
}

func (s *NestedConjunctionSearcher) toAdvanceID(key index.AncestorID) index.IndexInternalID {
	_ = "STUB: not implemented"
	return *new(index.IndexInternalID)
}

func (s *NestedConjunctionSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CoalesceQueue struct {
	order []*search.DocumentMatch
}

func NewCoalesceQueue() *CoalesceQueue { _ = "STUB: not implemented"; return nil }

func (cq *CoalesceQueue) Enqueue(it *search.DocumentMatch) { _ = "STUB: not implemented"; return }

func (cq *CoalesceQueue) Finalize() { _ = "STUB: not implemented"; return }

func (cq *CoalesceQueue) Dequeue(ctx *search.SearchContext) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (cq *CoalesceQueue) Len() int { _ = "STUB: not implemented"; return 0 }
