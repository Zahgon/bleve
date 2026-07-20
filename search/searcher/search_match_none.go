package searcher

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeMatchNoneSearcher int

func init() {
	var mns MatchNoneSearcher
	reflectStaticSizeMatchNoneSearcher = int(reflect.TypeOf(mns).Size())
}

type MatchNoneSearcher struct {
	indexReader index.IndexReader
}

func NewMatchNoneSearcher(indexReader index.IndexReader) (*MatchNoneSearcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MatchNoneSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *MatchNoneSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *MatchNoneSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *MatchNoneSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *MatchNoneSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MatchNoneSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MatchNoneSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (s *MatchNoneSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *MatchNoneSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
