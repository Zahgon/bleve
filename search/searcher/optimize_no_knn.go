//go:build !vectors
// +build !vectors

package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func optimizeKNN(ctx context.Context, indexReader index.IndexReader,
	qsearchers []search.Searcher) error {
	_ = "STUB: not implemented"
	return nil
}
