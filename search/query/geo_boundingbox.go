package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type GeoBoundingBoxQuery struct {
	TopLeft     []float64 `json:"top_left,omitempty"`
	BottomRight []float64 `json:"bottom_right,omitempty"`
	FieldVal    string    `json:"field,omitempty"`
	BoostVal    *Boost    `json:"boost,omitempty"`
}

func NewGeoBoundingBoxQuery(topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64) *GeoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *GeoBoundingBoxQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *GeoBoundingBoxQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *GeoBoundingBoxQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *GeoBoundingBoxQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *GeoBoundingBoxQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *GeoBoundingBoxQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *GeoBoundingBoxQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
