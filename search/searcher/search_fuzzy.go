package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var MaxFuzziness = 2

var AutoFuzzinessHighThreshold = 5

var AutoFuzzinessLowThreshold = 2

func NewFuzzySearcher(ctx context.Context, indexReader index.IndexReader, term string,
	prefix, fuzziness int, field string, boost float64,
	options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func GetAutoFuzziness(term string) int { _ = "STUB: not implemented"; return 0 }

func NewAutoFuzzySearcher(ctx context.Context, indexReader index.IndexReader, term string,
	prefix int, field string, boost float64, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

type fuzzyCandidates struct {
	candidates    []string
	editDistances []uint8
	bytesRead     uint64
}

func reportIOStats(ctx context.Context, bytesRead uint64) { _ = "STUB: not implemented"; return }

func findFuzzyCandidateTerms(ctx context.Context, indexReader index.IndexReader, term string,
	fuzziness int, field, prefixTerm string) (rv *fuzzyCandidates, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
