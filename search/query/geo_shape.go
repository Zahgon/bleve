package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type Geometry struct {
	Shape    index.GeoJSON `json:"shape"`
	Relation string        `json:"relation"`
}

type GeoShapeQuery struct {
	Geometry Geometry `json:"geometry"`
	FieldVal string   `json:"field,omitempty"`
	BoostVal *Boost   `json:"boost,omitempty"`
}

func NewGeoShapeQuery(coordinates [][][][]float64, typ,
	relation string) (*GeoShapeQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeoShapeCircleQuery(coordinates []float64, radius,
	relation string) (*GeoShapeQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeometryCollectionQuery(coordinates [][][][][]float64, types []string,
	relation string) (*GeoShapeQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *GeoShapeQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *GeoShapeQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *GeoShapeQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *GeoShapeQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *GeoShapeQuery) Searcher(ctx context.Context, i index.IndexReader,
	m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *GeoShapeQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *Geometry) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
