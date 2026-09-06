package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var DisjunctionMaxClauseCount = 0

var DisjunctionHeapTakeover = 10

func NewDisjunctionSearcher(ctx context.Context, indexReader index.IndexReader,
	qsearchers []search.Searcher, min float64, options search.SearcherOptions) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func optionsDisjunctionOptimizable(options search.SearcherOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func newDisjunctionSearcher(ctx context.Context, indexReader index.IndexReader,
	qsearchers []search.Searcher, min float64, options search.SearcherOptions,
	limit bool) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func optimizeCompositeSearcher(ctx context.Context, optimizationKind string,
	indexReader index.IndexReader, qsearchers []search.Searcher,
	options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func tooManyClauses(count int) bool { _ = "STUB: not implemented"; return false }

func tooManyClausesErr(field string, count int) error { _ = "STUB: not implemented"; return nil }
