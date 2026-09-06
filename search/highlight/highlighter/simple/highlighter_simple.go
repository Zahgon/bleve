package simple

import (
	index "github.com/blevesearch/bleve_index_api"

	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/highlight"
)

const Name = "simple"
const DefaultSeparator = "…"

type Highlighter struct {
	fragmenter highlight.Fragmenter
	formatter  highlight.FragmentFormatter
	sep        string
}

func NewHighlighter(fragmenter highlight.Fragmenter, formatter highlight.FragmentFormatter, separator string) *Highlighter {
	_ = "STUB: not implemented"
	return nil
}

func (s *Highlighter) Fragmenter() highlight.Fragmenter {
	_ = "STUB: not implemented"
	return *new(highlight.Fragmenter)
}

func (s *Highlighter) SetFragmenter(f highlight.Fragmenter) { _ = "STUB: not implemented"; return }

func (s *Highlighter) FragmentFormatter() highlight.FragmentFormatter {
	_ = "STUB: not implemented"
	return *new(highlight.FragmentFormatter)
}

func (s *Highlighter) SetFragmentFormatter(f highlight.FragmentFormatter) {
	_ = "STUB: not implemented"
	return
}

func (s *Highlighter) Separator() string { _ = "STUB: not implemented"; return "" }

func (s *Highlighter) SetSeparator(sep string) { _ = "STUB: not implemented"; return }

func (s *Highlighter) BestFragmentInField(dm *search.DocumentMatch, doc index.Document, field string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *Highlighter) BestFragmentsInField(dm *search.DocumentMatch, doc index.Document, field string, num int) []string {
	_ = "STUB: not implemented"
	return nil
}

type FragmentQueue []*highlight.Fragment

func (fq FragmentQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (fq FragmentQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (fq FragmentQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (fq *FragmentQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (fq *FragmentQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func Constructor(config map[string]interface{}, cache *registry.Cache) (highlight.Highlighter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Highlighter), nil
}

func init() {
	err := registry.RegisterHighlighter(Name, Constructor)
	if err != nil {
		panic(err)
	}
}
