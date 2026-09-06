package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeDocIDSearcher int

func init() {
	var ds DocIDSearcher
	reflectStaticSizeDocIDSearcher = int(reflect.TypeOf(ds).Size())
}

type DocIDSearcher struct {
	reader index.DocIDReader
	scorer *scorer.ConstantScorer
	count  int
}

func NewDocIDSearcher(ctx context.Context, indexReader index.IndexReader, ids []string, boost float64,
	options search.SearcherOptions) (searcher *DocIDSearcher, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DocIDSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *DocIDSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *DocIDSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *DocIDSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *DocIDSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DocIDSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *DocIDSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (s *DocIDSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *DocIDSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
