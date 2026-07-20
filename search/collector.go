package search

import (
	"context"
	"time"

	index "github.com/blevesearch/bleve_index_api"
)

type Collector interface {
	Collect(ctx context.Context, searcher Searcher, reader index.IndexReader) error
	Results() DocumentMatchCollection
	Total() uint64
	MaxScore() float64
	Took() time.Duration
	SetFacetsBuilder(facetsBuilder *FacetsBuilder)
	FacetResults() FacetResults
}

type DocumentMatchHandler func(hit *DocumentMatch) error

type MakeDocumentMatchHandlerKeyType string

var MakeDocumentMatchHandlerKey = MakeDocumentMatchHandlerKeyType(
	"MakeDocumentMatchHandlerKey")

var MakeKNNDocumentMatchHandlerKey = MakeDocumentMatchHandlerKeyType(
	"MakeKNNDocumentMatchHandlerKey")

type MakeDocumentMatchHandler func(ctx *SearchContext) (
	callback DocumentMatchHandler, loadID bool, err error)

type MakeKNNDocumentMatchHandler func(ctx *SearchContext) (
	callback DocumentMatchHandler, err error)
