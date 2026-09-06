package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeMatchAllSearcher int

func init() {
	var mas MatchAllSearcher
	reflectStaticSizeMatchAllSearcher = int(reflect.TypeOf(mas).Size())
}

type MatchAllSearcher struct {
	indexReader index.IndexReader
	reader      index.DocIDReader
	scorer      *scorer.ConstantScorer
	count       uint64
}

func NewMatchAllSearcher(ctx context.Context, indexReader index.IndexReader, boost float64, options search.SearcherOptions) (*MatchAllSearcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MatchAllSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *MatchAllSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *MatchAllSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *MatchAllSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *MatchAllSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MatchAllSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MatchAllSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (s *MatchAllSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *MatchAllSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
