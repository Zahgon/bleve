package searcher

import (
	"github.com/blevesearch/bleve/v2/search"
)

type OrderedSearcherList []search.Searcher

func (otrl OrderedSearcherList) Len() int { _ = "STUB: not implemented"; return 0 }

func (otrl OrderedSearcherList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (otrl OrderedSearcherList) Swap(i, j int) { _ = "STUB: not implemented"; return }

type OrderedPositionalSearcherList struct {
	searchers []search.Searcher
	index     []int
}

func (otrl OrderedPositionalSearcherList) Len() int { _ = "STUB: not implemented"; return 0 }

func (otrl OrderedPositionalSearcherList) Less(i, j int) bool {
	_ = "STUB: not implemented"
	return false
}

func (otrl OrderedPositionalSearcherList) Swap(i, j int) { _ = "STUB: not implemented"; return }
