package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/geo"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type GeoBoundingPolygonQuery struct {
	Points   []geo.Point `json:"polygon_points"`
	FieldVal string      `json:"field,omitempty"`
	BoostVal *Boost      `json:"boost,omitempty"`
}

func NewGeoBoundingPolygonQuery(points []geo.Point) *GeoBoundingPolygonQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *GeoBoundingPolygonQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *GeoBoundingPolygonQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *GeoBoundingPolygonQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *GeoBoundingPolygonQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *GeoBoundingPolygonQuery) Searcher(ctx context.Context, i index.IndexReader,
	m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *GeoBoundingPolygonQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *GeoBoundingPolygonQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
