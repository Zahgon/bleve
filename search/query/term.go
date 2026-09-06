package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type TermQuery struct {
	Term     string `json:"term"`
	FieldVal string `json:"field,omitempty"`
	BoostVal *Boost `json:"boost,omitempty"`
}

func NewTermQuery(term string) *TermQuery { _ = "STUB: not implemented"; return nil }

func (q *TermQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *TermQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *TermQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *TermQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *TermQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}
