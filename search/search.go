package search

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var (
	reflectStaticSizeDocumentMatch int
	reflectStaticSizeSearchContext int
	reflectStaticSizeLocation      int
)

func init() {
	var dm DocumentMatch
	reflectStaticSizeDocumentMatch = int(reflect.TypeOf(dm).Size())
	var sc SearchContext
	reflectStaticSizeSearchContext = int(reflect.TypeOf(sc).Size())
	var l Location
	reflectStaticSizeLocation = int(reflect.TypeOf(l).Size())
}

type ArrayPositions []uint64

func (ap ArrayPositions) Equals(other ArrayPositions) bool { _ = "STUB: not implemented"; return false }

func (ap ArrayPositions) Compare(other ArrayPositions) int { _ = "STUB: not implemented"; return 0 }

type Location struct {
	Pos uint64 `json:"pos"`

	Start uint64 `json:"start"`
	End   uint64 `json:"end"`

	ArrayPositions ArrayPositions `json:"array_positions"`
}

func (l *Location) Size() int { _ = "STUB: not implemented"; return 0 }

type Locations []*Location

func (p Locations) Len() int      { _ = "STUB: not implemented"; return 0 }
func (p Locations) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (p Locations) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (p Locations) Dedupe() Locations { _ = "STUB: not implemented"; return *new(Locations) }

type TermLocationMap map[string]Locations

func (t TermLocationMap) AddLocation(term string, location *Location) {
	_ = "STUB: not implemented"
	return
}

type FieldTermLocationMap map[string]TermLocationMap

type FieldTermLocation struct {
	Field    string
	Term     string
	Location Location
}

type FieldFragmentMap map[string][]string

type DocumentMatch struct {
	Index           string                `json:"index,omitempty"`
	ID              string                `json:"id"`
	IndexInternalID index.IndexInternalID `json:"-"`
	Score           float64               `json:"score"`
	Expl            *Explanation          `json:"explanation,omitempty"`
	Locations       FieldTermLocationMap  `json:"locations,omitempty"`
	Fragments       FieldFragmentMap      `json:"fragments,omitempty"`
	Sort            []string              `json:"sort,omitempty"`
	DecodedSort     []string              `json:"decoded_sort,omitempty"`

	Fields map[string]interface{} `json:"fields,omitempty"`

	HitNumber uint64 `json:"-"`

	FieldTermLocations []FieldTermLocation `json:"-"`

	ScoreBreakdown map[int]float64 `json:"score_breakdown,omitempty"`

	IndexNames []string `json:"index_names,omitempty"`

	Descendants []index.IndexInternalID `json:"-"`
}

func (dm *DocumentMatch) AddFieldValue(name string, value interface{}) {
	_ = "STUB: not implemented"
	return
}

func (dm *DocumentMatch) AddFragments(field string, fragments []string) {
	_ = "STUB: not implemented"
	return
}

func (dm *DocumentMatch) Reset() *DocumentMatch { _ = "STUB: not implemented"; return nil }

func (dm *DocumentMatch) Size() int { _ = "STUB: not implemented"; return 0 }

func (dm *DocumentMatch) Complete(prealloc []Location) []Location {
	_ = "STUB: not implemented"
	return nil
}

func (dm *DocumentMatch) String() string { _ = "STUB: not implemented"; return "" }

type DocumentMatchCollection []*DocumentMatch

func (c DocumentMatchCollection) Len() int           { _ = "STUB: not implemented"; return 0 }
func (c DocumentMatchCollection) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (c DocumentMatchCollection) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type Searcher interface {
	Next(ctx *SearchContext) (*DocumentMatch, error)
	Advance(ctx *SearchContext, ID index.IndexInternalID) (*DocumentMatch, error)
	Close() error
	Weight() float64
	SetQueryNorm(float64)
	Count() uint64
	Min() int
	Size() int

	DocumentMatchPoolSize() int
}

type SearcherOptions struct {
	Explain            bool
	IncludeTermVectors bool
	Score              string
}

type SearchContext struct {
	DocumentMatchPool *DocumentMatchPool
	Collector         Collector
	IndexReader       index.IndexReader
}

func (sc *SearchContext) Size() int { _ = "STUB: not implemented"; return 0 }

type NestedDocumentMatch struct {
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Fragments FieldFragmentMap       `json:"fragments,omitempty"`
}

func NewNestedDocumentMatch(fields map[string]interface{}, fragments FieldFragmentMap) *NestedDocumentMatch {
	_ = "STUB: not implemented"
	return nil
}
