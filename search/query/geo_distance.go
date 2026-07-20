package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type GeoDistanceQuery struct {
	Location []float64 `json:"location,omitempty"`
	Distance string    `json:"distance,omitempty"`
	FieldVal string    `json:"field,omitempty"`
	BoostVal *Boost    `json:"boost,omitempty"`
}

func NewGeoDistanceQuery(lon, lat float64, distance string) *GeoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *GeoDistanceQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *GeoDistanceQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *GeoDistanceQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *GeoDistanceQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *GeoDistanceQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping,
	options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *GeoDistanceQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *GeoDistanceQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
