//go:build vectors
// +build vectors

package searcher

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/scorer"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeKNNSearcher int

func init() {
	var ks KNNSearcher
	reflectStaticSizeKNNSearcher = int(reflect.TypeOf(ks).Size())
}

type KNNSearcher struct {
	field        string
	vector       []float32
	k            int64
	indexReader  index.IndexReader
	vectorReader index.VectorReader
	scorer       *scorer.KNNQueryScorer
	count        uint64
	vd           index.VectorDoc
}

func NewKNNSearcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping,
	options search.SearcherOptions, field string, vector []float32, k int64,
	boost float64, similarityMetric string, searchParams json.RawMessage,
	eligibleSelector index.EligibleDocumentSelector) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (s *KNNSearcher) VectorOptimize(ctx context.Context, octx index.VectorOptimizableContext) (
	index.VectorOptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.VectorOptimizableContext), nil
}

func (s *KNNSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (
	*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *KNNSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (s *KNNSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *KNNSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }

func (s *KNNSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *KNNSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *KNNSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *KNNSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *KNNSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }
