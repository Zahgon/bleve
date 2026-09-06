package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type DisjunctionQuery struct {
	Disjuncts              []Query `json:"disjuncts"`
	BoostVal               *Boost  `json:"boost,omitempty"`
	Min                    float64 `json:"min"`
	retrieveScoreBreakdown bool
	queryStringMode        bool
}

func (q *DisjunctionQuery) RetrieveScoreBreakdown(b bool) { _ = "STUB: not implemented"; return }

func NewDisjunctionQuery(disjuncts []Query) *DisjunctionQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *DisjunctionQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *DisjunctionQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *DisjunctionQuery) AddQuery(aq ...Query) { _ = "STUB: not implemented"; return }

func (q *DisjunctionQuery) SetMin(m float64) { _ = "STUB: not implemented"; return }

func (q *DisjunctionQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping,
	options search.SearcherOptions,
) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *DisjunctionQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *DisjunctionQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
