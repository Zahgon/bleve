package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/searcher"
	index "github.com/blevesearch/bleve_index_api"
)

type CustomFilterQuery struct {
	Query  Query    `json:"query"`
	Fields []string `json:"fields,omitempty"`

	filterFunc searcher.CustomFilterFunc
	payload    map[string]interface{}
}

var CustomFilterQueryParser func([]byte) (Query, error)

func NewCustomFilterQueryWithFilter(query Query, filter searcher.CustomFilterFunc, fields []string, payload map[string]interface{}) *CustomFilterQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *CustomFilterQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *CustomFilterQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *CustomFilterQuery) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *CustomFilterQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
