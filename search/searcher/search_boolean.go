package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeBooleanSearcher int

func init() {
	var bs BooleanSearcher
	reflectStaticSizeBooleanSearcher = int(reflect.TypeOf(bs).Size())
}

type BooleanSearcher struct {
	indexReader     index.IndexReader
	mustSearcher    search.Searcher
	shouldSearcher  search.Searcher
	mustNotSearcher search.Searcher
	queryNorm       float64
	currMust        *search.DocumentMatch
	currShould      *search.DocumentMatch
	currMustNot     *search.DocumentMatch
	currentID       index.IndexInternalID
	min             uint64
	scorer          *scorer.ConjunctionQueryScorer
	matches         []*search.DocumentMatch
	initialized     bool
	done            bool
}

func NewBooleanSearcher(ctx context.Context, indexReader index.IndexReader, mustSearcher search.Searcher, shouldSearcher search.Searcher, mustNotSearcher search.Searcher, options search.SearcherOptions) (*BooleanSearcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *BooleanSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *BooleanSearcher) computeQueryNorm() { _ = "STUB: not implemented"; return }

func (s *BooleanSearcher) initSearchers(ctx *search.SearchContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *BooleanSearcher) advanceNextMust(ctx *search.SearchContext, skipReturn *search.DocumentMatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *BooleanSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *BooleanSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *BooleanSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *BooleanSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *BooleanSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *BooleanSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (s *BooleanSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *BooleanSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
