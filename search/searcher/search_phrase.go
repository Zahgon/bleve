package searcher

import (
	"context"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizePhraseSearcher int

func init() {
	var ps PhraseSearcher
	reflectStaticSizePhraseSearcher = int(reflect.TypeOf(ps).Size())
}

type PhraseSearcher struct {
	mustSearcher search.Searcher
	queryNorm    float64
	currMust     *search.DocumentMatch
	terms        [][]string
	path         phrasePath
	paths        []phrasePath
	locations    []search.Location
	initialized  bool

	fuzzyTermMatches map[string][]string
}

func (s *PhraseSearcher) Size() int { _ = "STUB: not implemented"; return 0 }

func NewPhraseSearcher(ctx context.Context, indexReader index.IndexReader, terms []string,
	fuzziness int, autoFuzzy bool, field string, boost float64, options search.SearcherOptions) (*PhraseSearcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMultiPhraseSearcher(ctx context.Context, indexReader index.IndexReader, terms [][]string,
	fuzziness int, autoFuzzy bool, field string, boost float64, options search.SearcherOptions) (*PhraseSearcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PhraseSearcher) computeQueryNorm() { _ = "STUB: not implemented"; return }

func (s *PhraseSearcher) initSearchers(ctx *search.SearchContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PhraseSearcher) advanceNextMust(ctx *search.SearchContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PhraseSearcher) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *PhraseSearcher) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *PhraseSearcher) Next(ctx *search.SearchContext) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PhraseSearcher) checkCurrMustMatch(ctx *search.SearchContext) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *PhraseSearcher) checkCurrMustMatchField(ctx *search.SearchContext,
	field string, tlm search.TermLocationMap,
	ftls []search.FieldTermLocation) []search.FieldTermLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *PhraseSearcher) expandFuzzyMatches(tlm search.TermLocationMap, expandedTlm search.TermLocationMap) {
	_ = "STUB: not implemented"
	return
}

type phrasePart struct {
	term string
	loc  *search.Location
}

func (p *phrasePart) String() string { _ = "STUB: not implemented"; return "" }

type phrasePath []phrasePart

func (p phrasePath) MergeInto(in search.TermLocationMap) { _ = "STUB: not implemented"; return }

func (p phrasePath) String() string { _ = "STUB: not implemented"; return "" }

func findPhrasePaths(prevPos uint64, ap search.ArrayPositions, phraseTerms [][]string,
	tlm search.TermLocationMap, p phrasePath, remainingSlop int, rv []phrasePath) []phrasePath {
	_ = "STUB: not implemented"
	return nil
}

func editDistance(p1, p2 uint64) int { _ = "STUB: not implemented"; return 0 }

func (s *PhraseSearcher) Advance(ctx *search.SearchContext, ID index.IndexInternalID) (*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PhraseSearcher) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *PhraseSearcher) Close() error { _ = "STUB: not implemented"; return nil }

func (s *PhraseSearcher) Min() int { _ = "STUB: not implemented"; return 0 }

func (s *PhraseSearcher) DocumentMatchPoolSize() int { _ = "STUB: not implemented"; return 0 }
