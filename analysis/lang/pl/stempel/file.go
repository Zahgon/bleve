package stempel

import (
	_ "embed"
	"io"
)

//go:embed pl/stemmer_20000.tbl
var stempelFile []byte

type Trie interface {
	GetLastOnPath([]rune) []rune
}

func Open(path string) (Trie, error) { _ = "STUB: not implemented"; return *new(Trie), nil }

func LoadTrie() (Trie, error) { _ = "STUB: not implemented"; return *new(Trie), nil }

func buildTrieFromReader(f io.Reader) (Trie, error) {
	_ = "STUB: not implemented"
	return *new(Trie), nil
}
