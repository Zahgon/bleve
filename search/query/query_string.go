package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type QueryStringQuery struct {
	Query    string `json:"query"`
	BoostVal *Boost `json:"boost,omitempty"`
}

func NewQueryStringQuery(query string) *QueryStringQuery { _ = "STUB: not implemented"; return nil }

func (q *QueryStringQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *QueryStringQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *QueryStringQuery) Parse() (Query, error) {
	_ = "STUB: not implemented"
	return *new(Query), nil
}

func (q *QueryStringQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *QueryStringQuery) Validate() error { _ = "STUB: not implemented"; return nil }
