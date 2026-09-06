package stempel

import (
	"fmt"

	"github.com/blevesearch/stempel/javadata"
)

type multiTrie struct {
	tries   []*trie
	by      int32
	forward bool
}

func newMultiTrie(r *javadata.Reader) (rv *multiTrie, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const eom = rune('*')

func (t *multiTrie) GetLastOnPath(key []rune) []rune { _ = "STUB: not implemented"; return nil }

func cannotFollow(after, goes rune) bool { _ = "STUB: not implemented"; return false }

var errIndexOutOfBounds = fmt.Errorf("index out of bounds")

func (t *multiTrie) skip(in []rune, count int) ([]rune, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lengthPP(cmd []rune) int { _ = "STUB: not implemented"; return 0 }

func (t *multiTrie) String() string { _ = "STUB: not implemented"; return "" }
