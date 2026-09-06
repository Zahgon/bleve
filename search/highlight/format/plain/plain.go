package plain

import (
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search/highlight"
)

const Name = "plain"

const defaultPlainHighlightBefore = "<start>"
const defaultPlainHighlightAfter = "<end>"

type FragmentFormatter struct {
	before string
	after  string
}

func NewFragmentFormatter(before, after string) *FragmentFormatter {
	_ = "STUB: not implemented"
	return nil
}

func (a *FragmentFormatter) Format(f *highlight.Fragment, orderedTermLocations highlight.TermLocations) string {
	_ = "STUB: not implemented"
	return ""
}

func Constructor(config map[string]interface{}, cache *registry.Cache) (highlight.FragmentFormatter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.FragmentFormatter), nil
}

func init() {
	err := registry.RegisterFragmentFormatter(Name, Constructor)
	if err != nil {
		panic(err)
	}
}
