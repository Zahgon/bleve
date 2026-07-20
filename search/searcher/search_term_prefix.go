package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func NewTermPrefixSearcher(ctx context.Context, indexReader index.IndexReader, prefix string,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}
