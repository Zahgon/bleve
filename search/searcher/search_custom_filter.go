package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeCustomFilterSearcher int

func init() {
	var cfs CustomFilterSearcher
	reflectStaticSizeCustomFilterSearcher = int(reflect.TypeOf(cfs).Size())
}

type CustomFilterFunc func(ctx context.Context, d *search.DocumentMatch) (bool, error)

type CustomFilterSearcher struct {
	ctx         context.Context
	child       search.Searcher
	accept      CustomFilterFunc
	dvReader    index.DocValueReader
	indexReader index.IndexReader
	fieldTypes  map[string]string
}

func NewCustomFilterSearcher(ctx context.Context, child search.Searcher,
	filter CustomFilterFunc, dvReader index.DocValueReader,
	indexReader index.IndexReader,
	fieldTypes map[string]string) *CustomFilterSearcher {
	_ = "STUB: not implemented"
	return nil
}

func (f *CustomFilterSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (f *CustomFilterSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *CustomFilterSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *CustomFilterSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (f *CustomFilterSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (f *CustomFilterSearcher) SetQueryNorm(n float64) { _ = "STUB: not implemented"; return }

func (f *CustomFilterSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *CustomFilterSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (f *CustomFilterSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
