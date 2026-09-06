package query

import (
	"context"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type MatchPhraseQuery struct {
	MatchPhrase string `json:"match_phrase"`
	FieldVal    string `json:"field,omitempty"`
	Analyzer    string `json:"analyzer,omitempty"`
	BoostVal    *Boost `json:"boost,omitempty"`
	Fuzziness   int    `json:"fuzziness"`
	autoFuzzy   bool
}

func NewMatchPhraseQuery(matchPhrase string) *MatchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *MatchPhraseQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *MatchPhraseQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *MatchPhraseQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *MatchPhraseQuery) SetFuzziness(f int) { _ = "STUB: not implemented"; return }

func (q *MatchPhraseQuery) SetAutoFuzziness(auto bool) { _ = "STUB: not implemented"; return }

func (q *MatchPhraseQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *MatchPhraseQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func tokenStreamToPhrase(tokens analysis.TokenStream) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func (q *MatchPhraseQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (f *MatchPhraseQuery) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
