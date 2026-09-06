package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type IPRangeQuery struct {
	CIDR     string `json:"cidr,omitempty"`
	FieldVal string `json:"field,omitempty"`
	BoostVal *Boost `json:"boost,omitempty"`
}

func NewIPRangeQuery(cidr string) *IPRangeQuery { _ = "STUB: not implemented"; return nil }

func (q *IPRangeQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *IPRangeQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *IPRangeQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *IPRangeQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *IPRangeQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *IPRangeQuery) Validate() error { _ = "STUB: not implemented"; return nil }
