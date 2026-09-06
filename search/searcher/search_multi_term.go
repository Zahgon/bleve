package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func NewMultiTermSearcher(ctx context.Context, indexReader index.IndexReader, terms []string,
	field string, boost float64, options search.SearcherOptions, limit bool) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func NewMultiTermSearcherBoosted(ctx context.Context, indexReader index.IndexReader, terms []string,
	field string, boost float64, editDistances []uint8, options search.SearcherOptions, limit bool) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func NewMultiTermSearcherBytes(ctx context.Context, indexReader index.IndexReader, terms [][]byte,
	field string, boost float64, options search.SearcherOptions, limit bool) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func newMultiTermSearcherInternal(ctx context.Context, indexReader index.IndexReader,
	searchers []search.Searcher, field string, boost float64,
	options search.SearcherOptions, limit bool) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func optimizeMultiTermSearcher(ctx context.Context, indexReader index.IndexReader, terms []string,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func makeBatchSearchers(ctx context.Context, indexReader index.IndexReader, terms []string, field string,
	boost float64, options search.SearcherOptions) ([]search.Searcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeBatchSearchersBoosted(ctx context.Context, indexReader index.IndexReader, terms []string, field string,
	boost float64, editDistances []uint8, options search.SearcherOptions) ([]search.Searcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func optimizeMultiTermSearcherBytes(ctx context.Context, indexReader index.IndexReader, terms [][]byte,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func makeBatchSearchersBytes(ctx context.Context, indexReader index.IndexReader, terms [][]byte, field string,
	boost float64, options search.SearcherOptions) ([]search.Searcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
