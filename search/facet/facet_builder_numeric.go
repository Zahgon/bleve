package facet

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
)

var (
	reflectStaticSizeNumericFacetBuilder int
	reflectStaticSizenumericRange        int
)

func init() {
	var nfb NumericFacetBuilder
	reflectStaticSizeNumericFacetBuilder = int(reflect.TypeOf(nfb).Size())
	var nr numericRange
	reflectStaticSizenumericRange = int(reflect.TypeOf(nr).Size())
}

type numericRange struct {
	min *float64
	max *float64
}

type NumericFacetBuilder struct {
	size       int
	field      string
	termsCount map[string]int
	total      int
	missing    int
	ranges     map[string]*numericRange
	sawValue   bool
}

func NewNumericFacetBuilder(field string, size int) *NumericFacetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (fb *NumericFacetBuilder) Size() int { _ = "STUB: not implemented"; return 0 }

func (fb *NumericFacetBuilder) AddRange(name string, min, max *float64) {
	_ = "STUB: not implemented"
	return
}

func (fb *NumericFacetBuilder) Field() string { _ = "STUB: not implemented"; return "" }

func (fb *NumericFacetBuilder) UpdateVisitor(term []byte) { _ = "STUB: not implemented"; return }

func (fb *NumericFacetBuilder) StartDoc() { _ = "STUB: not implemented"; return }

func (fb *NumericFacetBuilder) EndDoc() { _ = "STUB: not implemented"; return }

func (fb *NumericFacetBuilder) Result() *search.FacetResult { _ = "STUB: not implemented"; return nil }
