package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type Regexp interface {
	FindStringIndex(s string) (loc []int)

	LiteralPrefix() (prefix string, complete bool)

	String() string
}

func NewRegexpStringSearcher(ctx context.Context, indexReader index.IndexReader, pattern string,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func NewRegexpSearcher(ctx context.Context, indexReader index.IndexReader, pattern Regexp,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

type regexpCandidates struct {
	candidates []string
	bytesRead  uint64
}

func findRegexpCandidateTerms(indexReader index.IndexReader,
	pattern Regexp, field, prefixTerm string) (rv *regexpCandidates, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
