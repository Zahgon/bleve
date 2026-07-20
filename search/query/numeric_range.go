package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type NumericRangeQuery struct {
	Min          *float64 `json:"min,omitempty"`
	Max          *float64 `json:"max,omitempty"`
	InclusiveMin *bool    `json:"inclusive_min,omitempty"`
	InclusiveMax *bool    `json:"inclusive_max,omitempty"`
	FieldVal     string   `json:"field,omitempty"`
	BoostVal     *Boost   `json:"boost,omitempty"`
}

func NewNumericRangeQuery(min, max *float64) *NumericRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewNumericRangeInclusiveQuery(min, max *float64, minInclusive, maxInclusive *bool) *NumericRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *NumericRangeQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *NumericRangeQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *NumericRangeQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *NumericRangeQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *NumericRangeQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *NumericRangeQuery) Validate() error { _ = "STUB: not implemented"; return nil }
