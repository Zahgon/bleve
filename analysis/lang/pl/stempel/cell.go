package stempel

import (
	"github.com/blevesearch/stempel/javadata"
)

type cell struct {
	ref int32
	cmd int32
}

func (c *cell) String() string { _ = "STUB: not implemented"; return "" }

func newCell(r *javadata.Reader) (*cell, error) { _ = "STUB: not implemented"; return nil, nil }
