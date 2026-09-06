package collector

import (
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type collectStoreNested struct {
	descAdder search.DescendantAdderCallbackFn

	nr index.NestedReader

	currRoot *search.DocumentMatch

	currRootAncestorID index.AncestorID

	ancestors []index.AncestorID
}

func newStoreNested(nr index.NestedReader, descAdder search.DescendantAdderCallbackFn) *collectStoreNested {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectStoreNested) ProcessNestedDocument(ctx *search.SearchContext, doc *search.DocumentMatch) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collectStoreNested) Current() *search.DocumentMatch { _ = "STUB: not implemented"; return nil }
