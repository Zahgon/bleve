package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/searcher"
	index "github.com/blevesearch/bleve_index_api"
)

type CustomScoreQuery struct {
	Query  Query    `json:"query"`
	Fields []string `json:"fields,omitempty"`

	scoreFunc searcher.CustomScoreFunc
	payload   map[string]interface{}
}

var CustomScoreQueryParser func([]byte) (Query, error)

func NewCustomScoreQueryWithScorer(query Query, score searcher.CustomScoreFunc, fields []string, payload map[string]interface{}) *CustomScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *CustomScoreQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *CustomScoreQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func (q *CustomScoreQuery) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *CustomScoreQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
