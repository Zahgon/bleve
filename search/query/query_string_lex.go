package query

import (
	"bufio"
	"io"
	"strings"
	"sync"
)

var queryStringLexPool = sync.Pool{
	New: func() interface{} {
		return &queryStringLex{
			in: bufio.NewReader(strings.NewReader("")),
		}
	},
}

func getQueryStringLex(in io.Reader) *queryStringLex { _ = "STUB: not implemented"; return nil }

func putQueryStringLex(l *queryStringLex) { _ = "STUB: not implemented"; return }

const reservedChars = "+-=&|><!(){}[]^\"~*?:\\/ "

func unescape(escaped string) string { _ = "STUB: not implemented"; return "" }

type queryStringLex struct {
	in            *bufio.Reader
	buf           string
	currState     lexState
	currConsumed  bool
	inEscape      bool
	nextToken     *yySymType
	nextTokenType int
	seenDot       bool
	nextRune      rune
	nextRuneSize  int
	atEOF         bool
}

func (l *queryStringLex) reset() { _ = "STUB: not implemented"; return }

func (l *queryStringLex) Error(msg string) { _ = "STUB: not implemented"; return }

func (l *queryStringLex) Lex(lval *yySymType) int { _ = "STUB: not implemented"; return 0 }

type lexState func(l *queryStringLex, next rune, eof bool) (lexState, bool)

func startState(l *queryStringLex, next rune, eof bool) (lexState, bool) {
	_ = "STUB: not implemented"
	return *new(lexState), false
}

func inPhraseState(l *queryStringLex, next rune, eof bool) (lexState, bool) {
	_ = "STUB: not implemented"
	return *new(lexState), false
}

func singleCharOpState(l *queryStringLex, next rune, eof bool) (lexState, bool) {
	_ = "STUB: not implemented"
	return *new(lexState), false
}

func inBoostState(l *queryStringLex, next rune, eof bool) (lexState, bool) {
	_ = "STUB: not implemented"
	return *new(lexState), false
}

func inTildeState(l *queryStringLex, next rune, eof bool) (lexState, bool) {
	_ = "STUB: not implemented"
	return *new(lexState), false
}

func inNumOrStrState(l *queryStringLex, next rune, eof bool) (lexState, bool) {
	_ = "STUB: not implemented"
	return *new(lexState), false
}

func inStrState(l *queryStringLex, next rune, eof bool) (lexState, bool) {
	_ = "STUB: not implemented"
	return *new(lexState), false
}

func logDebugTokens(format string, v ...interface{}) { _ = "STUB: not implemented"; return }
