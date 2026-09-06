package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type GeoShapeV2Query struct {
	GeometryV2 Geometry `json:"geometry_v2,omitempty"`
	FieldVal   string   `json:"field,omitempty"`
	BoostVal   *Boost   `json:"boost,omitempty"`
}

func NewGeoShapeV2Query(coordinates [][][][]float64, typ,
	relation string) (*GeoShapeV2Query, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeoShapeV2CircleQuery(center []float64, radius,
	relation string) (*GeoShapeV2Query, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeoShapeV2GeometryCollectionQuery(coordinates [][][][][]float64,
	types []string, relation string) (*GeoShapeV2Query, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *GeoShapeV2Query) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *GeoShapeV2Query) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *GeoShapeV2Query) Field() string { _ = "STUB: not implemented"; return "" }

func (q *GeoShapeV2Query) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *GeoShapeV2Query) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *GeoShapeV2Query) Searcher(ctx context.Context,
	i index.IndexReader, m mapping.IndexMapping,
	options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}
