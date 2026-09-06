package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type BooleanQuery struct {
	Must            Query  `json:"must,omitempty"`
	Should          Query  `json:"should,omitempty"`
	MustNot         Query  `json:"must_not,omitempty"`
	Filter          Query  `json:"filter,omitempty"`
	BoostVal        *Boost `json:"boost,omitempty"`
	queryStringMode bool
}

func NewBooleanQuery(must []Query, should []Query, mustNot []Query) *BooleanQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewBooleanQueryForQueryString(must []Query, should []Query, mustNot []Query) *BooleanQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *BooleanQuery) SetMinShould(minShould float64) { _ = "STUB: not implemented"; return }

func (q *BooleanQuery) AddMust(m ...Query) { _ = "STUB: not implemented"; return }

func (q *BooleanQuery) AddShould(m ...Query) { _ = "STUB: not implemented"; return }

func (q *BooleanQuery) AddMustNot(m ...Query) { _ = "STUB: not implemented"; return }

func (q *BooleanQuery) AddFilter(m Query) { _ = "STUB: not implemented"; return }

func (q *BooleanQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *BooleanQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *BooleanQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *BooleanQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *BooleanQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
