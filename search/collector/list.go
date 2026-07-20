package collector

import (
	"container/list"

	"github.com/blevesearch/bleve/v2/search"
)

type collectStoreList struct {
	results *list.List
	compare collectorCompare
}

func newStoreList(capacity int, compare collectorCompare) *collectStoreList {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreList) AddNotExceedingSize(doc *search.DocumentMatch, size int) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreList) add(doc *search.DocumentMatch) { _ = "STUB: not implemented"; return }

func (c *collectStoreList) removeLast() *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreList) Final(skip int, fixup collectorFixup) (search.DocumentMatchCollection, error) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection), nil
}

func (c *collectStoreList) Internal() search.DocumentMatchCollection {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection)
}

func (c *collectStoreList) len() int { _ = "STUB: not implemented"; return 0 }
