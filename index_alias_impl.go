package bleve

import (
	"context"
	"sync"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type indexAliasImpl struct {
	name    string
	indexes []Index
	mutex   sync.RWMutex
	open    bool

	mapping mapping.IndexMapping
}

func NewIndexAlias(indexes ...Index) *indexAliasImpl { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) VisitIndexes(visit func(Index)) { _ = "STUB: not implemented"; return }

func (i *indexAliasImpl) isAliasToSingleIndex() error { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) Index(id string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexAliasImpl) IndexSynonym(id string, collection string, definition *SynonymDefinition) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexAliasImpl) Train(batch *Batch) error { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) Delete(id string) error { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) Batch(b *Batch) error { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) Document(id string) (index.Document, error) {
	_ = "STUB: not implemented"
	return *new(index.Document), nil
}

func (i *indexAliasImpl) DocCount() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *indexAliasImpl) Search(req *SearchRequest) (*SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexAliasImpl) SearchInContext(ctx context.Context, req *SearchRequest) (*SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexAliasImpl) Fields() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *indexAliasImpl) FieldDict(field string) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *indexAliasImpl) FieldDictRange(field string, startTerm []byte, endTerm []byte) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *indexAliasImpl) FieldDictPrefix(field string, termPrefix []byte) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *indexAliasImpl) Close() error { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) SetIndexMapping(m mapping.IndexMapping) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexAliasImpl) Mapping() mapping.IndexMapping {
	_ = "STUB: not implemented"
	return *new(mapping.IndexMapping)
}

func (i *indexAliasImpl) Stats() *IndexStat { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) StatsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) GetInternal(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexAliasImpl) SetInternal(key, val []byte) error { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) DeleteInternal(key []byte) error { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) Advanced() (index.Index, error) {
	_ = "STUB: not implemented"
	return *new(index.Index), nil
}

func (i *indexAliasImpl) Add(indexes ...Index) { _ = "STUB: not implemented"; return }

func (i *indexAliasImpl) removeSingle(index Index) { _ = "STUB: not implemented"; return }

func (i *indexAliasImpl) Remove(indexes ...Index) { _ = "STUB: not implemented"; return }

func (i *indexAliasImpl) Swap(in, out []Index) { _ = "STUB: not implemented"; return }

func createChildSearchRequest(req *SearchRequest, preSearchData map[string]interface{}) *SearchRequest {
	_ = "STUB: not implemented"
	return nil
}

type asyncSearchResult struct {
	Name   string
	Result *SearchResult
	Err    error
}

type preSearchFlags struct {
	knn      bool
	synonyms bool
	bm25     bool
}

func isBM25Enabled(m mapping.IndexMapping) bool { _ = "STUB: not implemented"; return false }

func preSearchRequired(ctx context.Context, req *SearchRequest, m mapping.IndexMapping) (*preSearchFlags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func preSearch(ctx context.Context, req *SearchRequest, flags *preSearchFlags, indexes ...Index) (*SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finalizeSearchResult(ctx context.Context, req *SearchRequest, preSearchResult *SearchResult, rescorer *rescorer) *SearchResult {
	_ = "STUB: not implemented"
	return nil
}

func requestSatisfiedByPreSearch(req *SearchRequest, flags *preSearchFlags) bool {
	_ = "STUB: not implemented"
	return false
}

func constructSynonymPreSearchData(rv map[string]map[string]interface{}, sr *SearchResult, indexes []Index) map[string]map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func constructBM25PreSearchData(rv map[string]map[string]interface{}, sr *SearchResult, indexes []Index) map[string]map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func constructPreSearchData(req *SearchRequest, flags *preSearchFlags,
	preSearchResult *SearchResult, indexes []Index,
) (map[string]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func constructPreSearchDataAndFusionKnnHits(req *SearchRequest, flags *preSearchFlags,
	preSearchResult *SearchResult, rescorer *rescorer, indexes []Index,
) (map[string]map[string]interface{}, search.DocumentMatchCollection, error) {
	_ = "STUB: not implemented"
	return nil, *new(search.DocumentMatchCollection), nil
}

func preSearchDataSearch(ctx context.Context, req *SearchRequest, flags *preSearchFlags, indexes ...Index) (*SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func redistributePreSearchData(req *SearchRequest, indexes []Index) (map[string]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finalizePreSearchResult(req *SearchRequest, flags *preSearchFlags, preSearchResult *SearchResult) {
	_ = "STUB: not implemented"
	return
}

func hitsInCurrentPage(req *SearchRequest, hits []*search.DocumentMatch) []*search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

type multiSearchParams struct {
	preSearchData map[string]map[string]interface{}
	rescorer      *rescorer
	fusionKnnHits search.DocumentMatchCollection
}

func MultiSearch(ctx context.Context, req *SearchRequest, params *multiSearchParams, indexes ...Index) (*SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexAliasImpl) NewBatch() *Batch { _ = "STUB: not implemented"; return nil }

func (i *indexAliasImpl) Name() string { _ = "STUB: not implemented"; return "" }

func (i *indexAliasImpl) SetName(name string) { _ = "STUB: not implemented"; return }

type indexAliasImplFieldDict struct {
	index     *indexAliasImpl
	fieldDict index.FieldDict
}

func (f *indexAliasImplFieldDict) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *indexAliasImplFieldDict) Next() (*index.DictEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *indexAliasImplFieldDict) Close() error { _ = "STUB: not implemented"; return nil }

func (f *indexAliasImplFieldDict) Cardinality() int { _ = "STUB: not implemented"; return 0 }

func (i *indexAliasImpl) TermFrequencies(field string, limit int, descending bool) (
	[]index.TermFreq, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *indexAliasImpl) CentroidCardinalities(field string, limit int, descending bool) (
	[]index.CentroidCardinality, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
