package collector

import (
	"context"
	"reflect"
	"time"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeTopNCollector int

func init() {
	var coll TopNCollector
	reflectStaticSizeTopNCollector = int(reflect.TypeOf(coll).Size())
}

type collectorStore interface {
	AddNotExceedingSize(doc *search.DocumentMatch, size int) *search.DocumentMatch

	Final(skip int, fixup collectorFixup) (search.DocumentMatchCollection, error)

	Internal() search.DocumentMatchCollection
}

var PreAllocSizeSkipCap = 1000

type collectorCompare func(i, j *search.DocumentMatch) int

type collectorFixup func(d *search.DocumentMatch) error

type TopNCollector struct {
	size          int
	skip          int
	total         uint64
	bytesRead     uint64
	maxScore      float64
	took          time.Duration
	sort          search.SortOrder
	results       search.DocumentMatchCollection
	facetsBuilder *search.FacetsBuilder

	store collectorStore
	cmp   collectorCompare

	needDocIds    bool
	neededFields  []string
	cachedScoring []bool
	cachedDesc    []bool

	lowestMatchOutsideResults *search.DocumentMatch
	updateFieldVisitor        index.DocValueVisitor
	dvReader                  index.DocValueReader
	searchAfter               *search.DocumentMatch

	knnHits             map[string]*search.DocumentMatch
	hybridMergeCallback search.HybridMergeCallbackFn

	nestedStore *collectStoreNested

	fastPrepare bool

	earlyStopN   int
	earlyStopped bool
}

const CheckDoneEvery = uint64(1024)

func NewTopNCollector(size int, skip int, sort search.SortOrder) *TopNCollector {
	_ = "STUB: not implemented"
	return nil
}

func NewTopNCollectorAfter(size int, sort search.SortOrder, after []string) *TopNCollector {
	_ = "STUB: not implemented"
	return nil
}

func NewNestedTopNCollector(size int, skip int, sort search.SortOrder, nr index.NestedReader) *TopNCollector {
	_ = "STUB: not implemented"
	return nil
}

func NewNestedTopNCollectorAfter(size int, sort search.SortOrder, after []string, nr index.NestedReader) *TopNCollector {
	_ = "STUB: not implemented"
	return nil
}

func newTopNCollector(size int, skip int, sort search.SortOrder, nr index.NestedReader) *TopNCollector {
	_ = "STUB: not implemented"
	return nil
}

func getOptimalCollectorCompare(hc *TopNCollector) collectorCompare {
	_ = "STUB: not implemented"
	return *new(collectorCompare)
}

func createSearchAfterDocument(sort search.SortOrder, after []string) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func encodeSearchAfter(ss search.SearchSort, after string) string {
	_ = "STUB: not implemented"
	return ""
}

func FilterHitsBySearchAfter(hits []*search.DocumentMatch, sort search.SortOrder, after []string) []*search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func getOptimalCollectorStore(size, skip int, comparator collectorCompare) collectorStore {
	_ = "STUB: not implemented"
	return *new(collectorStore)
}

func (hc *TopNCollector) Size() int { _ = "STUB: not implemented"; return 0 }

func (hc *TopNCollector) Collect(ctx context.Context, searcher search.Searcher, reader index.IndexReader) error {
	_ = "STUB: not implemented"
	return nil
}

var sortByScoreOpt = []string{"_score"}

func (hc *TopNCollector) adjustKNNDocumentMatch(ctx *search.SearchContext,
	reader index.IndexReader, d *search.DocumentMatch) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (hc *TopNCollector) basicPrepare(d *search.DocumentMatch) { _ = "STUB: not implemented"; return }

func (hc *TopNCollector) canFastPrepare() bool { _ = "STUB: not implemented"; return false }

func (hc *TopNCollector) prepareDocumentMatch(ctx *search.SearchContext,
	reader index.IndexReader, d *search.DocumentMatch) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (hc *TopNCollector) prepareKNNDocumentMatch(ctx *search.SearchContext,
	reader index.IndexReader, d *search.DocumentMatch) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func MakeTopNDocumentMatchHandler(
	ctx *search.SearchContext) (search.DocumentMatchHandler, bool, error) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchHandler), false, nil
}

func (hc *TopNCollector) visitFieldTerms(reader index.IndexReader, d *search.DocumentMatch, v index.DocValueVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *TopNCollector) SetFacetsBuilder(facetsBuilder *search.FacetsBuilder) {
	_ = "STUB: not implemented"
	return
}

func (hc *TopNCollector) finalizeResults(r index.IndexReader) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *TopNCollector) Results() search.DocumentMatchCollection {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection)
}

func (hc *TopNCollector) Total() uint64 { _ = "STUB: not implemented"; return 0 }

func (hc *TopNCollector) SetEarlyStop(n int) { _ = "STUB: not implemented"; return }

func (hc *TopNCollector) EarlyStopped() bool { _ = "STUB: not implemented"; return false }

func (hc *TopNCollector) MaxScore() float64 { _ = "STUB: not implemented"; return 0 }

func (hc *TopNCollector) Took() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (hc *TopNCollector) FacetResults() search.FacetResults {
	_ = "STUB: not implemented"
	return *new(search.FacetResults)
}

func (hc *TopNCollector) SetKNNHits(knnHits search.DocumentMatchCollection, hybridMergeCallback search.HybridMergeCallbackFn) {
	_ = "STUB: not implemented"
	return
}
