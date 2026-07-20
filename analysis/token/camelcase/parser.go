package camelcase

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func (p *Parser) buildTokenFromTerm(buffer []rune) *analysis.Token {
	_ = "STUB: not implemented"
	return nil
}

type Parser struct {
	bufferLen int
	buffer    []rune
	current   State
	tokens    []*analysis.Token
	position  int
	index     int
}

func NewParser(length, position, index int) *Parser { _ = "STUB: not implemented"; return nil }

func (p *Parser) Push(sym rune, peek *rune) { _ = "STUB: not implemented"; return }

func (p *Parser) NewState(sym rune) State { _ = "STUB: not implemented"; return *new(State) }

func (p *Parser) FlushTokens() []*analysis.Token { _ = "STUB: not implemented"; return nil }

func (p *Parser) NextPosition() int { _ = "STUB: not implemented"; return 0 }
