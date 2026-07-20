//go:build vectors
// +build vectors

package collector

import (
	"context"
	"time"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type collectStoreKNN struct {
	internalHeaps []collectorStore
	kValues       []int64
	allHits       map[*search.DocumentMatch]struct{}
	ejectedDocs   map[*search.DocumentMatch]struct{}
}

func newStoreKNN(internalHeaps []collectorStore, kValues []int64) *collectStoreKNN {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreKNN) AddDocument(doc *search.DocumentMatch) []*search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreKNN) Final(fixup collectorFixup) (search.DocumentMatchCollection, error) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection), nil
}

func MakeKNNDocMatchHandler(ctx *search.SearchContext) (search.DocumentMatchHandler, error) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchHandler), nil
}

func GetNewKNNCollectorStore(kArray []int64) *collectStoreKNN {
	_ = "STUB: not implemented"
	return nil
}

type KNNCollector struct {
	knnStore *collectStoreKNN
	size     int
	total    uint64
	took     time.Duration
	results  search.DocumentMatchCollection
	maxScore float64
}

func NewKNNCollector(kArray []int64, size int64) *KNNCollector {
	_ = "STUB: not implemented"
	return nil
}

func (hc *KNNCollector) Collect(ctx context.Context, searcher search.Searcher, reader index.IndexReader) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *KNNCollector) finalizeResults(r index.IndexReader) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *KNNCollector) Results() search.DocumentMatchCollection {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection)
}

func (hc *KNNCollector) Total() uint64 { _ = "STUB: not implemented"; return 0 }

func (hc *KNNCollector) MaxScore() float64 { _ = "STUB: not implemented"; return 0 }

func (hc *KNNCollector) Took() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (hc *KNNCollector) SetFacetsBuilder(facetsBuilder *search.FacetsBuilder) {
	_ = "STUB: not implemented"
	return
}

func (hc *KNNCollector) FacetResults() search.FacetResults {
	_ = "STUB: not implemented"
	return *new(search.FacetResults)
}
