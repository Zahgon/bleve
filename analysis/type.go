package analysis

import (
	"fmt"
	"time"
)

type CharFilter interface {
	Filter([]byte) []byte
}

type TokenType int

const (
	AlphaNumeric TokenType = iota
	Ideographic
	Numeric
	DateTime
	Shingle
	Single
	Double
	Boolean
	IP
)

type Token struct {
	Start int `json:"start"`

	End  int    `json:"end"`
	Term []byte `json:"term"`

	Position int       `json:"position"`
	Type     TokenType `json:"type"`
	KeyWord  bool      `json:"keyword"`
}

func (t *Token) String() string { _ = "STUB: not implemented"; return "" }

type TokenStream []*Token

type Tokenizer interface {
	Tokenize([]byte) TokenStream
}

type TokenFilter interface {
	Filter(TokenStream) TokenStream
}

type Analyzer interface {
	Analyze([]byte) TokenStream
}

type DefaultAnalyzer struct {
	CharFilters  []CharFilter
	Tokenizer    Tokenizer
	TokenFilters []TokenFilter
}

func (a *DefaultAnalyzer) Analyze(input []byte) TokenStream {
	_ = "STUB: not implemented"
	return *new(TokenStream)
}

var ErrInvalidDateTime = fmt.Errorf("unable to parse datetime with any of the layouts")

var ErrInvalidTimestampString = fmt.Errorf("unable to parse timestamp string")
var ErrInvalidTimestampRange = fmt.Errorf("timestamp out of range")

type DateTimeParser interface {
	ParseDateTime(string) (time.Time, string, error)
}

const SynonymSourceType = "synonym"

type SynonymSourceVisitor func(name string, item SynonymSource) error

type SynonymSource interface {
	Analyzer() string
	Collection() string
}

type ByteArrayConverter interface {
	Convert([]byte) (interface{}, error)
}
