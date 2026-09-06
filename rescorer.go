package bleve

import (
	"github.com/blevesearch/bleve/v2/search"
)

const (
	DefaultScoreRankConstant = 60
)

type rescorer struct {
	req *SearchRequest

	origFrom   int
	origSize   int
	origBoosts []float64

	restored bool
}

func (r *rescorer) prepareSearchRequest() error { _ = "STUB: not implemented"; return nil }

func (r *rescorer) restoreSearchRequest() { _ = "STUB: not implemented"; return }

func (r *rescorer) rescore(ftsHits, knnHits search.DocumentMatchCollection) (search.DocumentMatchCollection, uint64, float64) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection), 0, 0
}

func (r *rescorer) mergeDocs(ftsHits, knnHits search.DocumentMatchCollection) search.DocumentMatchCollection {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection)
}

func newRescorer(req *SearchRequest) *rescorer { _ = "STUB: not implemented"; return nil }
