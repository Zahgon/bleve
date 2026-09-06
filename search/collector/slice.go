package collector

import (
	"github.com/blevesearch/bleve/v2/search"
)

type collectStoreSlice struct {
	slice   search.DocumentMatchCollection
	compare collectorCompare
}

func newStoreSlice(capacity int, compare collectorCompare) *collectStoreSlice {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreSlice) AddNotExceedingSize(doc *search.DocumentMatch,
	size int) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreSlice) add(doc *search.DocumentMatch) { _ = "STUB: not implemented"; return }

func (c *collectStoreSlice) removeLast() *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreSlice) Final(skip int, fixup collectorFixup) (search.DocumentMatchCollection, error) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection), nil
}

func (c *collectStoreSlice) Internal() search.DocumentMatchCollection {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection)
}

func (c *collectStoreSlice) len() int { _ = "STUB: not implemented"; return 0 }
