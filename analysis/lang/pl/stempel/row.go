package stempel

import (
	"github.com/blevesearch/stempel/javadata"
)

type row struct {
	cells map[rune]*cell
}

func (r *row) String() string { _ = "STUB: not implemented"; return "" }

func newRow(r *javadata.Reader) (*row, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *row) getCmd(way rune) int32 { _ = "STUB: not implemented"; return 0 }

func (r *row) getRef(way rune) int32 { _ = "STUB: not implemented"; return 0 }

func (r *row) at(c rune) *cell { _ = "STUB: not implemented"; return nil }
