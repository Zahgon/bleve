package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type RegexpQuery struct {
	Regexp   string `json:"regexp"`
	FieldVal string `json:"field,omitempty"`
	BoostVal *Boost `json:"boost,omitempty"`
}

func NewRegexpQuery(regexp string) *RegexpQuery { _ = "STUB: not implemented"; return nil }

func (q *RegexpQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *RegexpQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *RegexpQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *RegexpQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *RegexpQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *RegexpQuery) Validate() error { _ = "STUB: not implemented"; return nil }
