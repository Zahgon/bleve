package collector

import (
	"github.com/blevesearch/bleve/v2/search"
)

type collectStoreHeap struct {
	heap    search.DocumentMatchCollection
	compare collectorCompare
}

func newStoreHeap(capacity int, compare collectorCompare) *collectStoreHeap {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreHeap) AddNotExceedingSize(doc *search.DocumentMatch,
	size int) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreHeap) add(doc *search.DocumentMatch) { _ = "STUB: not implemented"; return }

func (c *collectStoreHeap) removeLast() *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreHeap) siftUp(i int) { _ = "STUB: not implemented"; return }

func (c *collectStoreHeap) siftDown(i int) { _ = "STUB: not implemented"; return }

func (c *collectStoreHeap) Final(skip int, fixup collectorFixup) (search.DocumentMatchCollection, error) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection), nil
}

func (c *collectStoreHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *collectStoreHeap) Internal() search.DocumentMatchCollection {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection)
}
