package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type MatchAllQuery struct {
	BoostVal *Boost `json:"boost,omitempty"`
}

func NewMatchAllQuery() *MatchAllQuery { _ = "STUB: not implemented"; return nil }

func (q *MatchAllQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *MatchAllQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *MatchAllQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *MatchAllQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
