package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func NewNumericRangeSearcher(ctx context.Context, indexReader index.IndexReader,
	min *float64, max *float64, inclusiveMin, inclusiveMax *bool, field string,
	boost float64, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func filterCandidateTerms(indexReader index.IndexReader,
	terms [][]byte, field string) (rv [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type termRange struct {
	startTerm []byte
	endTerm   []byte
}

func (t *termRange) Enumerate(filter filterFunc) [][]byte { _ = "STUB: not implemented"; return nil }

func incrementBytes(in []byte) []byte { _ = "STUB: not implemented"; return nil }

type termRanges []*termRange

func (tr termRanges) Enumerate(filter filterFunc) [][]byte { _ = "STUB: not implemented"; return nil }

func splitInt64Range(minBound, maxBound int64, precisionStep uint) termRanges {
	_ = "STUB: not implemented"
	return *new(termRanges)
}

func newRange(minBound, maxBound int64, shift uint) *termRange {
	_ = "STUB: not implemented"
	return nil
}

func newRangeBytes(minBytes, maxBytes []byte) *termRange { _ = "STUB: not implemented"; return nil }
