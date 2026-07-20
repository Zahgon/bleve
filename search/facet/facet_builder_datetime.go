package facet

import (
	"reflect"
	"time"

	"github.com/blevesearch/bleve/v2/search"
)

var (
	reflectStaticSizeDateTimeFacetBuilder int
	reflectStaticSizedateTimeRange        int
)

func init() {
	var dtfb DateTimeFacetBuilder
	reflectStaticSizeDateTimeFacetBuilder = int(reflect.TypeOf(dtfb).Size())
	var dtr dateTimeRange
	reflectStaticSizedateTimeRange = int(reflect.TypeOf(dtr).Size())
}

type dateTimeRange struct {
	start time.Time
	end   time.Time
}

type DateTimeFacetBuilder struct {
	size       int
	field      string
	termsCount map[string]int
	total      int
	missing    int
	ranges     map[string]*dateTimeRange
	sawValue   bool
}

func NewDateTimeFacetBuilder(field string, size int) *DateTimeFacetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (fb *DateTimeFacetBuilder) Size() int { _ = "STUB: not implemented"; return 0 }

func (fb *DateTimeFacetBuilder) AddRange(name string, start, end time.Time) {
	_ = "STUB: not implemented"
	return
}

func (fb *DateTimeFacetBuilder) Field() string { _ = "STUB: not implemented"; return "" }

func (fb *DateTimeFacetBuilder) UpdateVisitor(term []byte) { _ = "STUB: not implemented"; return }

func (fb *DateTimeFacetBuilder) StartDoc() { _ = "STUB: not implemented"; return }

func (fb *DateTimeFacetBuilder) EndDoc() { _ = "STUB: not implemented"; return }

func (fb *DateTimeFacetBuilder) Result() *search.FacetResult { _ = "STUB: not implemented"; return nil }
