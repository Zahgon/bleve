package query

import (
	"context"
	"strings"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var wildcardRegexpReplacer = strings.NewReplacer(

	"+", `\+`,
	"(", `\(`,
	")", `\)`,
	"^", `\^`,
	"$", `\$`,
	".", `\.`,
	"{", `\{`,
	"}", `\}`,
	"[", `\[`,
	"]", `\]`,
	`|`, `\|`,
	`\`, `\\`,

	"*", ".*",
	"?", ".")

type WildcardQuery struct {
	Wildcard string `json:"wildcard"`
	FieldVal string `json:"field,omitempty"`
	BoostVal *Boost `json:"boost,omitempty"`
}

func NewWildcardQuery(wildcard string) *WildcardQuery { _ = "STUB: not implemented"; return nil }

func (q *WildcardQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *WildcardQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *WildcardQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *WildcardQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *WildcardQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *WildcardQuery) Validate() error { _ = "STUB: not implemented"; return nil }
