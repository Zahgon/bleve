package search

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeFacetsBuilder int
var reflectStaticSizeFacetResult int
var reflectStaticSizeTermFacet int
var reflectStaticSizeNumericRangeFacet int
var reflectStaticSizeDateRangeFacet int

func init() {
	var fb FacetsBuilder
	reflectStaticSizeFacetsBuilder = int(reflect.TypeOf(fb).Size())
	var fr FacetResult
	reflectStaticSizeFacetResult = int(reflect.TypeOf(fr).Size())
	var tf TermFacet
	reflectStaticSizeTermFacet = int(reflect.TypeOf(tf).Size())
	var nrf NumericRangeFacet
	reflectStaticSizeNumericRangeFacet = int(reflect.TypeOf(nrf).Size())
	var drf DateRangeFacet
	reflectStaticSizeDateRangeFacet = int(reflect.TypeOf(drf).Size())
}

type FacetBuilder interface {
	StartDoc()
	UpdateVisitor(term []byte)
	EndDoc()

	Result() *FacetResult
	Field() string

	Size() int
}

type FacetsBuilder struct {
	indexReader   index.IndexReader
	facetNames    []string
	facets        []FacetBuilder
	facetsByField map[string][]FacetBuilder
	fields        []string
}

func NewFacetsBuilder(indexReader index.IndexReader) *FacetsBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (fb *FacetsBuilder) Size() int { _ = "STUB: not implemented"; return 0 }

func (fb *FacetsBuilder) Add(name string, facetBuilder FacetBuilder) {
	_ = "STUB: not implemented"
	return
}

func (fb *FacetsBuilder) RequiredFields() []string { _ = "STUB: not implemented"; return nil }

func (fb *FacetsBuilder) StartDoc() { _ = "STUB: not implemented"; return }

func (fb *FacetsBuilder) EndDoc() { _ = "STUB: not implemented"; return }

func (fb *FacetsBuilder) UpdateVisitor(field string, term []byte) {
	_ = "STUB: not implemented"
	return
}

type TermFacet struct {
	Term  string `json:"term"`
	Count int    `json:"count"`
}

type TermFacets struct {
	termFacets []*TermFacet
	termLookup map[string]*TermFacet
}

func (tf *TermFacets) Terms() []*TermFacet { _ = "STUB: not implemented"; return nil }

func (tf *TermFacets) TrimToTopN(n int) { _ = "STUB: not implemented"; return }

func (tf *TermFacets) Add(termFacets ...*TermFacet) { _ = "STUB: not implemented"; return }

func (tf *TermFacets) Len() int { _ = "STUB: not implemented"; return 0 }

func (tf *TermFacets) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (tf *TermFacets) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (tf *TermFacets) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tf *TermFacets) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type NumericRangeFacet struct {
	Name  string   `json:"name"`
	Min   *float64 `json:"min,omitempty"`
	Max   *float64 `json:"max,omitempty"`
	Count int      `json:"count"`
}

func (nrf *NumericRangeFacet) Same(other *NumericRangeFacet) bool {
	_ = "STUB: not implemented"
	return false
}

type NumericRangeFacets []*NumericRangeFacet

func (nrf NumericRangeFacets) Add(numericRangeFacet *NumericRangeFacet) NumericRangeFacets {
	_ = "STUB: not implemented"
	return *new(NumericRangeFacets)
}

func (nrf NumericRangeFacets) Len() int           { _ = "STUB: not implemented"; return 0 }
func (nrf NumericRangeFacets) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (nrf NumericRangeFacets) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type DateRangeFacet struct {
	Name  string  `json:"name"`
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
	Count int     `json:"count"`
}

func (drf *DateRangeFacet) Same(other *DateRangeFacet) bool {
	_ = "STUB: not implemented"
	return false
}

type DateRangeFacets []*DateRangeFacet

func (drf DateRangeFacets) Add(dateRangeFacet *DateRangeFacet) DateRangeFacets {
	_ = "STUB: not implemented"
	return *new(DateRangeFacets)
}

func (drf DateRangeFacets) Len() int           { _ = "STUB: not implemented"; return 0 }
func (drf DateRangeFacets) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (drf DateRangeFacets) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type FacetResult struct {
	Field         string             `json:"field"`
	Total         int                `json:"total"`
	Missing       int                `json:"missing"`
	Other         int                `json:"other"`
	Terms         *TermFacets        `json:"terms,omitempty"`
	NumericRanges NumericRangeFacets `json:"numeric_ranges,omitempty"`
	DateRanges    DateRangeFacets    `json:"date_ranges,omitempty"`
}

func (fr *FacetResult) Size() int { _ = "STUB: not implemented"; return 0 }

func (fr *FacetResult) Merge(other *FacetResult) { _ = "STUB: not implemented"; return }

func (fr *FacetResult) Fixup(size int) { _ = "STUB: not implemented"; return }

type FacetResults map[string]*FacetResult

func (fr FacetResults) Merge(other FacetResults) { _ = "STUB: not implemented"; return }

func (fr FacetResults) Fixup(name string, size int) { _ = "STUB: not implemented"; return }

func (fb *FacetsBuilder) Results() FacetResults {
	_ = "STUB: not implemented"
	return *new(FacetResults)
}
