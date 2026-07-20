package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type MatchQuery struct {
	Match     string             `json:"match"`
	FieldVal  string             `json:"field,omitempty"`
	Analyzer  string             `json:"analyzer,omitempty"`
	BoostVal  *Boost             `json:"boost,omitempty"`
	Prefix    int                `json:"prefix_length"`
	Fuzziness int                `json:"fuzziness"`
	Operator  MatchQueryOperator `json:"operator,omitempty"`
	autoFuzzy bool
}

type MatchQueryOperator int

const (
	MatchQueryOperatorOr = MatchQueryOperator(0)

	MatchQueryOperatorAnd = MatchQueryOperator(1)
)

func (o MatchQueryOperator) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *MatchQueryOperator) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMatchQuery(match string) *MatchQuery { _ = "STUB: not implemented"; return nil }

func (q *MatchQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *MatchQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *MatchQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *MatchQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *MatchQuery) SetFuzziness(f int) { _ = "STUB: not implemented"; return }

func (q *MatchQuery) SetAutoFuzziness(auto bool) { _ = "STUB: not implemented"; return }

func (q *MatchQuery) SetPrefix(p int) { _ = "STUB: not implemented"; return }

func (q *MatchQuery) SetOperator(operator MatchQueryOperator) { _ = "STUB: not implemented"; return }

func (q *MatchQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *MatchQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (f *MatchQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
