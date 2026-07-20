package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type FuzzyQuery struct {
	Term      string `json:"term"`
	Prefix    int    `json:"prefix_length"`
	Fuzziness int    `json:"fuzziness"`
	FieldVal  string `json:"field,omitempty"`
	BoostVal  *Boost `json:"boost,omitempty"`
	autoFuzzy bool
}

func NewFuzzyQuery(term string) *FuzzyQuery { _ = "STUB: not implemented"; return nil }

func (q *FuzzyQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *FuzzyQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *FuzzyQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *FuzzyQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *FuzzyQuery) SetFuzziness(f int) { _ = "STUB: not implemented"; return }

func (q *FuzzyQuery) SetAutoFuzziness(a bool) { _ = "STUB: not implemented"; return }

func (q *FuzzyQuery) SetPrefix(p int) { _ = "STUB: not implemented"; return }

func (q *FuzzyQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *FuzzyQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (f *FuzzyQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
