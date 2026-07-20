package stempel

import (
	"github.com/blevesearch/stempel/javadata"
)

type trie struct {
	rows    []*row
	cmds    []string
	root    int32
	forward bool
}

func newTrie(r *javadata.Reader) (rv *trie, err error) { _ = "STUB: not implemented"; return nil, nil }

func (t *trie) getRow(i int) *row { _ = "STUB: not implemented"; return nil }

func (t *trie) GetLastOnPath(key []rune) []rune { _ = "STUB: not implemented"; return nil }

func (t *trie) String() string { _ = "STUB: not implemented"; return "" }
