package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type PhraseQuery struct {
	Terms     []string `json:"terms"`
	FieldVal  string   `json:"field,omitempty"`
	BoostVal  *Boost   `json:"boost,omitempty"`
	Fuzziness int      `json:"fuzziness"`
	autoFuzzy bool
}

func NewPhraseQuery(terms []string, field string) *PhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *PhraseQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *PhraseQuery) SetFuzziness(f int) { _ = "STUB: not implemented"; return }

func (q *PhraseQuery) SetAutoFuzziness(auto bool) { _ = "STUB: not implemented"; return }

func (q *PhraseQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *PhraseQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *PhraseQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *PhraseQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *PhraseQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *PhraseQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (f *PhraseQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
