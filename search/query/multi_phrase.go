package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type MultiPhraseQuery struct {
	Terms     [][]string `json:"terms"`
	FieldVal  string     `json:"field,omitempty"`
	BoostVal  *Boost     `json:"boost,omitempty"`
	Fuzziness int        `json:"fuzziness"`
	autoFuzzy bool
}

func NewMultiPhraseQuery(terms [][]string, field string) *MultiPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *MultiPhraseQuery) SetFuzziness(f int) { _ = "STUB: not implemented"; return }

func (q *MultiPhraseQuery) SetAutoFuzziness(auto bool) { _ = "STUB: not implemented"; return }

func (q *MultiPhraseQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *MultiPhraseQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *MultiPhraseQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *MultiPhraseQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *MultiPhraseQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *MultiPhraseQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *MultiPhraseQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (f *MultiPhraseQuery) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
