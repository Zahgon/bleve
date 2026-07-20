package bleve

import (
	"encoding/json"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/analysis/datetime/optional"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/query"
)

var (
	reflectStaticSizeSearchResult int
	reflectStaticSizeSearchStatus int
)

func init() {
	var sr SearchResult
	reflectStaticSizeSearchResult = int(reflect.TypeOf(sr).Size())
	var ss SearchStatus
	reflectStaticSizeSearchStatus = int(reflect.TypeOf(ss).Size())
}

var cache = registry.NewCache()

const defaultDateTimeParser = optional.Name

const (
	ScoreDefault = ""
	ScoreNone    = "none"
	ScoreRRF     = "rrf"
	ScoreRSF     = "rsf"
)

var AllowedFusionSort = search.SortOrder{&search.SortScore{Desc: true}}

type dateTimeRange struct {
	Name           string    `json:"name,omitempty"`
	Start          time.Time `json:"start,omitempty"`
	End            time.Time `json:"end,omitempty"`
	DateTimeParser string    `json:"datetime_parser,omitempty"`
	startString    *string
	endString      *string
}

func (dr *dateTimeRange) ParseDates(dateTimeParser analysis.DateTimeParser) (start, end time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), nil
}

func (dr *dateTimeRange) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (dr *dateTimeRange) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type numericRange struct {
	Name string   `json:"name,omitempty"`
	Min  *float64 `json:"min,omitempty"`
	Max  *float64 `json:"max,omitempty"`
}

type FacetRequest struct {
	Size           int              `json:"size"`
	Field          string           `json:"field"`
	TermPrefix     string           `json:"term_prefix,omitempty"`
	TermPattern    string           `json:"term_pattern,omitempty"`
	NumericRanges  []*numericRange  `json:"numeric_ranges,omitempty"`
	DateTimeRanges []*dateTimeRange `json:"date_ranges,omitempty"`

	compiledPattern *regexp.Regexp
}

func NewFacetRequest(field string, size int) *FacetRequest { _ = "STUB: not implemented"; return nil }

func (fr *FacetRequest) SetPrefixFilter(prefix string) { _ = "STUB: not implemented"; return }

func (fr *FacetRequest) SetRegexFilter(pattern string) { _ = "STUB: not implemented"; return }

func (fr *FacetRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (fr *FacetRequest) AddDateTimeRange(name string, start, end time.Time) {
	_ = "STUB: not implemented"
	return
}

func (fr *FacetRequest) AddDateTimeRangeString(name string, start, end *string) {
	_ = "STUB: not implemented"
	return
}

func (fr *FacetRequest) AddDateTimeRangeStringWithParser(name string, start, end *string, parser string) {
	_ = "STUB: not implemented"
	return
}

func (fr *FacetRequest) AddNumericRange(name string, min, max *float64) {
	_ = "STUB: not implemented"
	return
}

type FacetsRequest map[string]*FacetRequest

func (fr FacetsRequest) Validate() error { _ = "STUB: not implemented"; return nil }

type HighlightRequest struct {
	Style  *string  `json:"style"`
	Fields []string `json:"fields"`
}

func NewHighlight() *HighlightRequest { _ = "STUB: not implemented"; return nil }

func NewHighlightWithStyle(style string) *HighlightRequest { _ = "STUB: not implemented"; return nil }

func (h *HighlightRequest) AddField(field string) { _ = "STUB: not implemented"; return }

func (r *SearchRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (r *SearchRequest) validatePagination() error { _ = "STUB: not implemented"; return nil }

func (r *SearchRequest) AddFacet(facetName string, f *FacetRequest) {
	_ = "STUB: not implemented"
	return
}

func (r *SearchRequest) SortBy(order []string) { _ = "STUB: not implemented"; return }

func (r *SearchRequest) SortByCustom(order search.SortOrder) { _ = "STUB: not implemented"; return }

func (r *SearchRequest) SetSearchAfter(after []string) { _ = "STUB: not implemented"; return }

func (r *SearchRequest) SetSearchBefore(before []string) { _ = "STUB: not implemented"; return }

func (r *SearchRequest) AddParams(params RequestParams) { _ = "STUB: not implemented"; return }

func NewSearchRequest(q query.Query) *SearchRequest { _ = "STUB: not implemented"; return nil }

func NewSearchRequestOptions(q query.Query, size, from int, explain bool) *SearchRequest {
	_ = "STUB: not implemented"
	return nil
}

type IndexErrMap map[string]error

func (iem IndexErrMap) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (iem IndexErrMap) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type SearchStatus struct {
	Total      int         `json:"total"`
	Failed     int         `json:"failed"`
	Successful int         `json:"successful"`
	Errors     IndexErrMap `json:"errors,omitempty"`
}

func (ss *SearchStatus) Merge(other *SearchStatus) { _ = "STUB: not implemented"; return }

type SearchResult struct {
	Status   *SearchStatus                  `json:"status"`
	Request  *SearchRequest                 `json:"request,omitempty"`
	Hits     search.DocumentMatchCollection `json:"hits"`
	Total    uint64                         `json:"total_hits"`
	Cost     uint64                         `json:"cost"`
	MaxScore float64                        `json:"max_score"`
	Took     time.Duration                  `json:"took"`
	Facets   search.FacetResults            `json:"facets"`

	SynonymResult search.FieldTermSynonymMap `json:"synonym_result,omitempty"`

	BM25Stats *search.BM25Stats `json:"bm25_stats,omitempty"`
}

func (sr *SearchResult) Size() int { _ = "STUB: not implemented"; return 0 }

func (sr *SearchResult) String() string { _ = "STUB: not implemented"; return "" }

func formatHit(rv *strings.Builder, hit *search.DocumentMatch, hitNumber int) *strings.Builder {
	_ = "STUB: not implemented"
	return nil
}

func (sr *SearchResult) Merge(other *SearchResult) { _ = "STUB: not implemented"; return }

func MemoryNeededForSearchResult(req *SearchRequest) uint64 { _ = "STUB: not implemented"; return 0 }

func (r *SearchRequest) SetSortFunc(s func(sort.Interface)) { _ = "STUB: not implemented"; return }

func (r *SearchRequest) SortFunc() func(data sort.Interface) { _ = "STUB: not implemented"; return nil }

func isMatchNoneQuery(q query.Query) bool { _ = "STUB: not implemented"; return false }

func isMatchAllQuery(q query.Query) bool { _ = "STUB: not implemented"; return false }

func IsScoreFusionRequested(req *SearchRequest) bool { _ = "STUB: not implemented"; return false }

type RequestParams struct {
	ScoreRankConstant int `json:"score_rank_constant,omitempty"`
	ScoreWindowSize   int `json:"score_window_size,omitempty"`
}

func NewDefaultParams(from, size int) *RequestParams { _ = "STUB: not implemented"; return nil }

func (p *RequestParams) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (p *RequestParams) Validate(size int) error { _ = "STUB: not implemented"; return nil }

func ParseParams(r *SearchRequest, input []byte) (*RequestParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type OptionalRawMessage json.RawMessage

func (n *OptionalRawMessage) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n OptionalRawMessage) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
