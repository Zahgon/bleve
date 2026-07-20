package query

import (
	"context"
	"io"
	"log"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var logger = log.New(io.Discard, "bleve mapping ", log.LstdFlags)

func SetLog(l *log.Logger) { _ = "STUB: not implemented"; return }

type Query interface {
	Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping,
		options search.SearcherOptions) (search.Searcher, error)
}

type BoostableQuery interface {
	Query
	SetBoost(b float64)
	Boost() float64
}

type FieldableQuery interface {
	Query
	SetField(f string)
	Field() string
}

type ValidatableQuery interface {
	Query
	Validate() error
}

func ParsePreSearchData(input []byte) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseQuery(input []byte) (Query, error) { _ = "STUB: not implemented"; return *new(Query), nil }

func expandQuery(m mapping.IndexMapping, query Query) (Query, error) {
	_ = "STUB: not implemented"
	return *new(Query), nil
}

func DumpQuery(m mapping.IndexMapping, query Query) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ExtractFields(q Query, m mapping.IndexMapping, fs search.FieldSet) (search.FieldSet, error) {
	_ = "STUB: not implemented"
	return *new(search.FieldSet), nil
}

const (
	FuzzyMatchType = iota
	RegexpMatchType
	PrefixMatchType
)

func ExtractSynonyms(ctx context.Context, m mapping.SynonymMapping, r index.ThesaurusReader,
	query Query, rv search.FieldTermSynonymMap,
) (search.FieldTermSynonymMap, error) {
	_ = "STUB: not implemented"
	return *new(search.FieldTermSynonymMap), nil
}

func addSynonymsForTermWithMatchType(ctx context.Context, matchType int, src, field, term string, fuzziness, prefix int,
	r index.ThesaurusReader, rv search.FieldTermSynonymMap,
) (search.FieldTermSynonymMap, error) {
	_ = "STUB: not implemented"
	return *new(search.FieldTermSynonymMap), nil
}

func addSynonymsForTerm(ctx context.Context, src, field, term string,
	r index.ThesaurusReader, rv search.FieldTermSynonymMap,
) (search.FieldTermSynonymMap, error) {
	_ = "STUB: not implemented"
	return *new(search.FieldTermSynonymMap), nil
}
