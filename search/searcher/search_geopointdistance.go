package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func NewGeoPointDistanceSearcher(ctx context.Context, indexReader index.IndexReader, centerLon,
	centerLat, dist float64, field string, boost float64,
	options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func boxSearcher(ctx context.Context, indexReader index.IndexReader,
	topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64,
	field string, boost float64, options search.SearcherOptions, checkBoundaries bool) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func buildDistFilter(ctx context.Context, dvReader index.DocValueReader,
	centerLon, centerLat, maxDist float64) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}
