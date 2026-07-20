package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func NewGeoShapeSearcher(ctx context.Context, indexReader index.IndexReader, shape index.GeoJSON,
	relation string, field string, boost float64,
	options search.SearcherOptions,
) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func buildRelationFilterOnShapes(ctx context.Context, dvReader index.DocValueReader, field string,
	relation string, shape index.GeoJSON,
) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}
