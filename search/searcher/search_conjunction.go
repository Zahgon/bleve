package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeConjunctionSearcher int

func init() {
	var cs ConjunctionSearcher
	reflectStaticSizeConjunctionSearcher = int(reflect.TypeOf(cs).Size())
}

type ConjunctionSearcher struct {
	indexReader index.IndexReader
	searchers   []search.Searcher
	queryNorm   float64
	currs       []*search.DocumentMatch
	maxIDIdx    int
	scorer      *scorer.ConjunctionQueryScorer
	initialized bool
	options     search.SearcherOptions
	bytesRead   uint64
}

func NewConjunctionSearcher(ctx context.Context, indexReader index.IndexReader,
	qsearchers []search.Searcher, options search.SearcherOptions) (
	search.Searcher, error,
) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (s *ConjunctionSearcher) computeQueryNorm() { _ = "STUB: not implemented"; return }

func (s *ConjunctionSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *ConjunctionSearcher) initSearchers(ctx *search.SearchContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ConjunctionSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *ConjunctionSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *ConjunctionSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ConjunctionSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ConjunctionSearcher) advanceChild(ctx *search.SearchContext, i int, ID index.IndexInternalID) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *ConjunctionSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *ConjunctionSearcher) Close() (rv error) { _ = "STUB: not implemented"; return nil }

func (s *ConjunctionSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *ConjunctionSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
