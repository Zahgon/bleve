package highlight

import (
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type Fragment struct {
	Orig           []byte
	ArrayPositions []uint64
	Start          int
	End            int
	Score          float64
	Index          int
}

func (f *Fragment) Overlaps(other *Fragment) bool { _ = "STUB: not implemented"; return false }

type Fragmenter interface {
	Fragment([]byte, TermLocations) []*Fragment
}

type FragmentFormatter interface {
	Format(f *Fragment, orderedTermLocations TermLocations) string
}

type FragmentScorer interface {
	Score(f *Fragment) float64
}

type Highlighter interface {
	Fragmenter() Fragmenter
	SetFragmenter(Fragmenter)

	FragmentFormatter() FragmentFormatter
	SetFragmentFormatter(FragmentFormatter)

	Separator() string
	SetSeparator(string)

	BestFragmentInField(*search.DocumentMatch, index.Document, string) string
	BestFragmentsInField(*search.DocumentMatch, index.Document, string, int) []string
}
