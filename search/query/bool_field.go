package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type BoolFieldQuery struct {
	Bool     bool   `json:"bool"`
	FieldVal string `json:"field,omitempty"`
	BoostVal *Boost `json:"boost,omitempty"`
}

func NewBoolFieldQuery(val bool) *BoolFieldQuery { _ = "STUB: not implemented"; return nil }

func (q *BoolFieldQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *BoolFieldQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *BoolFieldQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *BoolFieldQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *BoolFieldQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}
