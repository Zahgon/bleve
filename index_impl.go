package bleve

import (
	"context"
	"io"
	"sync"

	"github.com/blevesearch/bleve/v2/document"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/collector"
	"github.com/blevesearch/bleve/v2/search/highlight"
	index "github.com/blevesearch/bleve_index_api"
)

type indexImpl struct {
	path  string
	name  string
	meta  *indexMeta
	i     index.Index
	m     mapping.IndexMapping
	mutex sync.RWMutex
	open  bool
	stats *IndexStat
}

const storePath = "store"

const (
	SearchQueryStartCallbackKey search.ContextKey = "_search_query_start_callback_key"
	SearchQueryEndCallbackKey   search.ContextKey = "_search_query_end_callback_key"
)

type (
	SearchQueryStartCallbackFn func(size uint64) error
	SearchQueryEndCallbackFn   func(size uint64) error
)

func indexStorePath(path string) string { _ = "STUB: not implemented"; return "" }

func newIndexUsing(path string, mapping mapping.IndexMapping, indexType string, kvstore string, kvconfig map[string]interface{}) (*indexImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openIndexUsing(path string, runtimeConfig map[string]interface{}) (rv *indexImpl, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) Advanced() (index.Index, error) {
	_ = "STUB: not implemented"
	return *new(index.Index), nil
}

func (i *indexImpl) Mapping() mapping.IndexMapping {
	_ = "STUB: not implemented"
	return *new(mapping.IndexMapping)
}

func (i *indexImpl) Index(id string, data interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexImpl) IndexSynonym(id string, collection string, definition *SynonymDefinition) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexImpl) Train(batch *Batch) error { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) IndexAdvanced(doc *document.Document) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexImpl) Delete(id string) (err error) { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) Batch(b *Batch) error { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) Document(id string) (doc index.Document, err error) {
	_ = "STUB: not implemented"
	return *new(index.Document), nil
}

func (i *indexImpl) DocCount() (count uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

func (i *indexImpl) Search(req *SearchRequest) (sr *SearchResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) FileWriterIDsInUse() (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) DropFileWriterIDs(ids map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	documentMatchEmptySize int
	searchContextEmptySize int
	facetResultEmptySize   int
	documentEmptySize      int
)

func init() {
	var dm search.DocumentMatch
	documentMatchEmptySize = dm.Size()

	var sc search.SearchContext
	searchContextEmptySize = sc.Size()

	var fr search.FacetResult
	facetResultEmptySize = fr.Size()

	var d document.Document
	documentEmptySize = d.Size()
}

func memNeededForSearch(req *SearchRequest,
	searcher search.Searcher,
	topnCollector *collector.TopNCollector,
) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (i *indexImpl) preSearch(ctx context.Context, req *SearchRequest, reader index.IndexReader) (*SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) SearchInContext(ctx context.Context, req *SearchRequest) (sr *SearchResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadAndHighlightFields(hit *search.DocumentMatch, req *SearchRequest,
	indexName string, r index.IndexReader,
	highlighter highlight.Highlighter,
) (error, uint64) {
	_ = "STUB: not implemented"
	return nil, 0
}

const NestedDocumentKey = "_$nested"

func LoadAndHighlightAllFields(
	root *search.DocumentMatch,
	req *SearchRequest,
	indexName string,
	r index.IndexReader,
	highlighter highlight.Highlighter,
) (error, uint64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (i *indexImpl) Fields() (fields []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) FieldDict(field string) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *indexImpl) FieldDictRange(field string, startTerm []byte, endTerm []byte) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *indexImpl) FieldDictPrefix(field string, termPrefix []byte) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *indexImpl) Close() error { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) Stats() *IndexStat { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) StatsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) GetInternal(key []byte) (val []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) SetInternal(key, val []byte) error { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) DeleteInternal(key []byte) error { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) NewBatch() *Batch { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) Name() string { _ = "STUB: not implemented"; return "" }

func (i *indexImpl) SetName(name string) { _ = "STUB: not implemented"; return }

type indexImplFieldDict struct {
	index       *indexImpl
	indexReader index.IndexReader
	fieldDict   index.FieldDict
}

func (f *indexImplFieldDict) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *indexImplFieldDict) Next() (*index.DictEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *indexImplFieldDict) Close() error { _ = "STUB: not implemented"; return nil }

func (f *indexImplFieldDict) Cardinality() int { _ = "STUB: not implemented"; return 0 }

func deDuplicate(fields []string) []string { _ = "STUB: not implemented"; return nil }

type searchHitSorter struct {
	hits          search.DocumentMatchCollection
	sort          search.SortOrder
	cachedScoring []bool
	cachedDesc    []bool
}

func newSearchHitSorter(sort search.SortOrder, hits search.DocumentMatchCollection) *searchHitSorter {
	_ = "STUB: not implemented"
	return nil
}

func (m *searchHitSorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (m *searchHitSorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (m *searchHitSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (i *indexImpl) CopyTo(d index.Directory) (err error) { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) CopyFile(file string, d index.IndexDirectory) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexImpl) SetPathInBolt(key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FileSystemDirectory) GetWriter(filePath string) (io.WriteCloser,
	error,
) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (i *indexImpl) FireIndexEvent() { _ = "STUB: not implemented"; return }

func (i *indexImpl) TermFrequencies(field string, limit int, descending bool) (
	[]index.TermFreq, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) CentroidCardinalities(field string, limit int, descending bool) (
	[]index.CentroidCardinality, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexImpl) buildTopNCollector(ctx context.Context, req *SearchRequest, reader index.IndexReader) (*collector.TopNCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
