package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeGeoShapeV2Searcher int

func init() {
	var gsv2s GeoShapeV2Searcher
	reflectStaticSizeGeoShapeV2Searcher = int(reflect.TypeOf(gsv2s).Size())
}

type GeoShapeV2Searcher struct {
	geoShapeIndexReader index.GeoShapeV2FieldReader
	scorer              *scorer.ConstantScorer

	gd index.GeoShapeV2FieldDoc
}

func NewGeoShapeV2Searcher(ctx context.Context, indexReader index.IndexReader,
	shape index.GeoJSON, relation string, field string, boost float64,
	options search.SearcherOptions,
) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (g *GeoShapeV2Searcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoShapeV2Searcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoShapeV2Searcher) Close() error { _ = "STUB: not implemented"; return nil }

func (g *GeoShapeV2Searcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (g *GeoShapeV2Searcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }

func (g *GeoShapeV2Searcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (g *GeoShapeV2Searcher) SetQueryNorm(n float64) { _ = "STUB: not implemented"; return }

func (g *GeoShapeV2Searcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (g *GeoShapeV2Searcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }
