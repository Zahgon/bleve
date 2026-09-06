package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type DocIDQuery struct {
	IDs      []string `json:"ids"`
	BoostVal *Boost   `json:"boost,omitempty"`
}

func NewDocIDQuery(ids []string) *DocIDQuery { _ = "STUB: not implemented"; return nil }

func (q *DocIDQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *DocIDQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *DocIDQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}
