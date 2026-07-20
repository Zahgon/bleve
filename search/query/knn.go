//go:build vectors
// +build vectors

package query

import (
	"context"
	"encoding/json"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type KNNQuery struct {
	VectorField string    `json:"field"`
	Vector      []float32 `json:"vector"`
	K           int64     `json:"k"`
	BoostVal    *Boost    `json:"boost,omitempty"`

	Params json.RawMessage `json:"params"`

	elegibleSelector index.EligibleDocumentSelector
}

func NewKNNQuery(vector []float32) *KNNQuery { _ = "STUB: not implemented"; return nil }

func (q *KNNQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *KNNQuery) SetK(k int64) { _ = "STUB: not implemented"; return }

func (q *KNNQuery) SetField(field string) { _ = "STUB: not implemented"; return }

func (q *KNNQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *KNNQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *KNNQuery) SetParams(params json.RawMessage) { _ = "STUB: not implemented"; return }

func (q *KNNQuery) SetEligibleSelector(eligibleSelector index.EligibleDocumentSelector) {
	_ = "STUB: not implemented"
	return
}

func (q *KNNQuery) Searcher(ctx context.Context, i index.IndexReader,
	m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}
