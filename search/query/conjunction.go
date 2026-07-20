package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type ConjunctionQuery struct {
	Conjuncts       []Query `json:"conjuncts"`
	BoostVal        *Boost  `json:"boost,omitempty"`
	queryStringMode bool
}

func NewConjunctionQuery(conjuncts []Query) *ConjunctionQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *ConjunctionQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *ConjunctionQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *ConjunctionQuery) AddQuery(aq ...Query) { _ = "STUB: not implemented"; return }

func (q *ConjunctionQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *ConjunctionQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *ConjunctionQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
