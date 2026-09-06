//go:generate goyacc -o query_string.y.go query_string.y
//go:generate sed -i.tmp -e 1d query_string.y.go
//go:generate rm query_string.y.go.tmp

package query

var debugParser bool
var debugLexer bool

func parseQuerySyntax(query string) (rq Query, err error) {
	_ = "STUB: not implemented"
	return *new(Query), nil
}

func doParse(lex *lexerWrapper) { _ = "STUB: not implemented"; return }

const (
	queryShould = iota
	queryMust
	queryMustNot
)

type lexerWrapper struct {
	lex   yyLexer
	errs  []string
	query *BooleanQuery
}

func newLexerWrapper(lex yyLexer) *lexerWrapper { _ = "STUB: not implemented"; return nil }

func (l *lexerWrapper) Lex(lval *yySymType) int { _ = "STUB: not implemented"; return 0 }

func (l *lexerWrapper) Error(s string) { _ = "STUB: not implemented"; return }
