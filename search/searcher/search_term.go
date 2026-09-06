package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeTermSearcher int

func init() {
	var ts TermSearcher
	reflectStaticSizeTermSearcher = int(reflect.TypeOf(ts).Size())
}

type TermSearcher struct {
	indexReader index.IndexReader
	reader      index.TermFieldReader
	scorer      *scorer.TermQueryScorer
	tfd         index.TermFieldDoc
}

func NewTermSearcher(ctx context.Context, indexReader index.IndexReader,
	term string, field string, boost float64, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func NewTermSearcherBytes(ctx context.Context, indexReader index.IndexReader,
	term []byte, field string, boost float64, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func tfIDFScoreMetrics(indexReader index.IndexReader) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func bm25ScoreMetrics(ctx context.Context, field string,
	indexReader index.IndexReader) (uint64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func newTermSearcherFromReader(ctx context.Context, indexReader index.IndexReader,
	reader index.TermFieldReader, term []byte, field string, boost float64,
	options search.SearcherOptions) (*TermSearcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSynonymSearcher(ctx context.Context, indexReader index.IndexReader, term []byte, synonyms []string, field string, boost float64, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (s *TermSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *TermSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *TermSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *TermSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *TermSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TermSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TermSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (s *TermSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *TermSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }

func (s *TermSearcher) Optimize(kind string, octx index.OptimizableContext) (
	index.OptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.OptimizableContext), nil
}

func isTermQuery(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
