package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/geo"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func NewGeoBoundedPolygonSearcher(ctx context.Context, indexReader index.IndexReader,
	coordinates []geo.Point, field string, boost float64,
	options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

const float64EqualityThreshold = 1e-6

func almostEqual(a, b float64) bool { _ = "STUB: not implemented"; return false }

func buildPolygonFilter(ctx context.Context, dvReader index.DocValueReader, field string,
	coordinates []geo.Point) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}
