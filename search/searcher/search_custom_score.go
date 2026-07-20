package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeCustomScoreSearcher int

func init() {
	var sfs CustomScoreSearcher
	reflectStaticSizeCustomScoreSearcher = int(reflect.TypeOf(sfs).Size())
}

type CustomScoreFunc func(ctx context.Context, d *search.DocumentMatch) (float64, error)

type CustomScoreSearcher struct {
	ctx         context.Context
	child       search.Searcher
	mutate      CustomScoreFunc
	dvReader    index.DocValueReader
	indexReader index.IndexReader
	fieldTypes  map[string]string
	explain     bool
}

func NewCustomScoreSearcher(ctx context.Context, s search.Searcher, mutate CustomScoreFunc,
	dvReader index.DocValueReader, indexReader index.IndexReader,
	fieldTypes map[string]string, explain bool) *CustomScoreSearcher {
	_ = "STUB: not implemented"
	return nil
}

func (f *CustomScoreSearcher) applyScore(d *search.DocumentMatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *CustomScoreSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (f *CustomScoreSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *CustomScoreSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *CustomScoreSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (f *CustomScoreSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (f *CustomScoreSearcher) SetQueryNorm(n float64) { _ = "STUB: not implemented"; return }

func (f *CustomScoreSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *CustomScoreSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (f *CustomScoreSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
