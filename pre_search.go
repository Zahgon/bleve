package bleve

import (
	"github.com/blevesearch/bleve/v2/search"
)

type preSearchResultProcessor interface {
	add(*SearchResult, string)

	finalize(*SearchResult)
}

type knnPreSearchResultProcessor struct {
	addFn      func(sr *SearchResult, indexName string)
	finalizeFn func(sr *SearchResult)
}

func (k *knnPreSearchResultProcessor) add(sr *SearchResult, indexName string) {
	_ = "STUB: not implemented"
	return
}

func (k *knnPreSearchResultProcessor) finalize(sr *SearchResult) { _ = "STUB: not implemented"; return }

type synonymPreSearchResultProcessor struct {
	finalizedFts search.FieldTermSynonymMap
}

func newSynonymPreSearchResultProcessor() *synonymPreSearchResultProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *synonymPreSearchResultProcessor) add(sr *SearchResult, indexName string) {
	_ = "STUB: not implemented"
	return
}

func (s *synonymPreSearchResultProcessor) finalize(sr *SearchResult) {
	_ = "STUB: not implemented"
	return
}

type bm25PreSearchResultProcessor struct {
	docCount         float64
	fieldCardinality map[string]int
}

func newBM25PreSearchResultProcessor() *bm25PreSearchResultProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (b *bm25PreSearchResultProcessor) add(sr *SearchResult, indexName string) {
	_ = "STUB: not implemented"
	return
}

func (b *bm25PreSearchResultProcessor) finalize(sr *SearchResult) {
	_ = "STUB: not implemented"
	return
}

type compositePreSearchResultProcessor struct {
	presearchResultProcessors []preSearchResultProcessor
}

func (m *compositePreSearchResultProcessor) add(sr *SearchResult, indexName string) {
	_ = "STUB: not implemented"
	return
}

func (m *compositePreSearchResultProcessor) finalize(sr *SearchResult) {
	_ = "STUB: not implemented"
	return
}

func createPreSearchResultProcessor(req *SearchRequest, flags *preSearchFlags) preSearchResultProcessor {
	_ = "STUB: not implemented"
	return *new(preSearchResultProcessor)
}
