package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeFilteringSearcher int

func init() {
	var fs FilteringSearcher
	reflectStaticSizeFilteringSearcher = int(reflect.TypeOf(fs).Size())
}

type FilterFunc func(sctx *search.SearchContext, d *search.DocumentMatch) bool

type FilteringSearcher struct {
	child  search.Searcher
	accept FilterFunc
}

func NewFilteringSearcher(ctx context.Context, s search.Searcher, filter FilterFunc) *FilteringSearcher {
	_ = "STUB: not implemented"
	return nil
}

func (f *FilteringSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (f *FilteringSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilteringSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilteringSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (f *FilteringSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (f *FilteringSearcher) SetQueryNorm(n float64) { _ = "STUB: not implemented"; return }

func (f *FilteringSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *FilteringSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (f *FilteringSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
