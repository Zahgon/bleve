//go:build gofuzz
// +build gofuzz

package stempel

var fuzzTrie Trie

func init() {
	var err error
	fuzzTrie, err = Open("pl/stemmer_20000.tbl")
	if err != nil {
		panic(err)
	}
}

func Fuzz(data []byte) int { _ = "STUB: not implemented"; return 0 }
