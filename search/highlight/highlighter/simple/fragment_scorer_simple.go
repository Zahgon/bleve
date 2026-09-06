package simple

import (
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/highlight"
)

type FragmentScorer struct {
	tlm search.TermLocationMap
}

func NewFragmentScorer(tlm search.TermLocationMap) *FragmentScorer {
	_ = "STUB: not implemented"
	return nil
}

func (s *FragmentScorer) Score(f *highlight.Fragment) { _ = "STUB: not implemented"; return }
