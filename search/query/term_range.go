package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type TermRangeQuery struct {
	Min          string `json:"min,omitempty"`
	Max          string `json:"max,omitempty"`
	InclusiveMin *bool  `json:"inclusive_min,omitempty"`
	InclusiveMax *bool  `json:"inclusive_max,omitempty"`
	FieldVal     string `json:"field,omitempty"`
	BoostVal     *Boost `json:"boost,omitempty"`
}

func NewTermRangeQuery(min, max string) *TermRangeQuery { _ = "STUB: not implemented"; return nil }

func NewTermRangeInclusiveQuery(min, max string, minInclusive, maxInclusive *bool) *TermRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *TermRangeQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *TermRangeQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *TermRangeQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *TermRangeQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *TermRangeQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *TermRangeQuery) Validate() error { _ = "STUB: not implemented"; return nil }
