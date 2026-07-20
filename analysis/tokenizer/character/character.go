package character

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

type IsTokenRune func(r rune) bool

type CharacterTokenizer struct {
	isTokenRun IsTokenRune
}

func NewCharacterTokenizer(f IsTokenRune) *CharacterTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (c *CharacterTokenizer) Tokenize(input []byte) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}
