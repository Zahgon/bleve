package highlight

import (
	"github.com/blevesearch/bleve/v2/search"
)

type TermLocation struct {
	Term           string
	ArrayPositions search.ArrayPositions
	Pos            int
	Start          int
	End            int
}

func (tl *TermLocation) Overlaps(other *TermLocation) bool { _ = "STUB: not implemented"; return false }

type TermLocations []*TermLocation

func (t TermLocations) Len() int           { _ = "STUB: not implemented"; return 0 }
func (t TermLocations) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (t TermLocations) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (t TermLocations) MergeOverlapping() { _ = "STUB: not implemented"; return }

func OrderTermLocations(tlm search.TermLocationMap) TermLocations {
	_ = "STUB: not implemented"
	return *new(TermLocations)
}
