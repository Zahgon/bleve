package facet

import (
	"reflect"
	"regexp"

	"github.com/blevesearch/bleve/v2/search"
)

var reflectStaticSizeTermsFacetBuilder int

func init() {
	var tfb TermsFacetBuilder
	reflectStaticSizeTermsFacetBuilder = int(reflect.TypeOf(tfb).Size())
}

type TermsFacetBuilder struct {
	size        int
	field       string
	prefixBytes []byte
	regex       *regexp.Regexp
	termsCount  map[string]int
	total       int
	missing     int
	sawValue    bool
}

func NewTermsFacetBuilder(field string, size int) *TermsFacetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (fb *TermsFacetBuilder) Size() int { _ = "STUB: not implemented"; return 0 }

func (fb *TermsFacetBuilder) Field() string { _ = "STUB: not implemented"; return "" }

func (fb *TermsFacetBuilder) SetPrefixFilter(prefix string) { _ = "STUB: not implemented"; return }

func (fb *TermsFacetBuilder) SetRegexFilter(regex *regexp.Regexp) {
	_ = "STUB: not implemented"
	return
}

func (fb *TermsFacetBuilder) UpdateVisitor(term []byte) { _ = "STUB: not implemented"; return }

func (fb *TermsFacetBuilder) StartDoc() { _ = "STUB: not implemented"; return }

func (fb *TermsFacetBuilder) EndDoc() { _ = "STUB: not implemented"; return }

func (fb *TermsFacetBuilder) Result() *search.FacetResult { _ = "STUB: not implemented"; return nil }
