//line query_string.y:2
package query

//line query_string.y:2

func logDebugGrammar(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

//line query_string.y:17
type yySymType struct {
	yys int
	s   string
	n   int
	f   float64
	q   Query
	pf  *float64
}

const tSTRING = 57346
const tPHRASE = 57347
const tPLUS = 57348
const tMINUS = 57349
const tCOLON = 57350
const tBOOST = 57351
const tNUMBER = 57352
const tGREATER = 57353
const tLESS = 57354
const tEQUAL = 57355
const tTILDE = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tSTRING",
	"tPHRASE",
	"tPLUS",
	"tMINUS",
	"tCOLON",
	"tBOOST",
	"tNUMBER",
	"tGREATER",
	"tLESS",
	"tEQUAL",
	"tTILDE",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 3,
	1, 3,
	-2, 5,
	-1, 9,
	8, 29,
	-2, 8,
	-1, 12,
	8, 28,
	-2, 12,
}

const yyPrivate = 57344

const yyLast = 43

var yyAct = [...]int{
	18, 17, 19, 24, 23, 15, 31, 22, 20, 21,
	30, 27, 23, 23, 3, 22, 22, 14, 29, 26,
	16, 25, 28, 35, 33, 23, 23, 32, 22, 22,
	34, 9, 12, 1, 5, 6, 2, 11, 4, 13,
	7, 8, 10,
}

var yyPact = [...]int{
	28, -1000, -1000, 28, 27, -1000, -1000, -1000, 8, -9,
	12, -1000, -1000, -1000, -1000, -1000, -3, -11, -1000, -1000,
	6, 5, -1000, -4, -1000, -1000, 19, -1000, -1000, 18,
	-1000, -1000, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]int{
	0, 0, 42, 41, 39, 38, 33, 36, 14,
}

var yyR1 = [...]int{
	0, 6, 7, 7, 8, 5, 5, 5, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 1, 1, 2, 2,
}

var yyR2 = [...]int{
	0, 1, 2, 1, 3, 0, 1, 1, 1, 2,
	4, 1, 1, 3, 3, 3, 4, 5, 4, 5,
	4, 5, 4, 5, 0, 1, 1, 2, 1, 1,
}

var yyChk = [...]int{
	-1000, -6, -7, -8, -5, 6, 7, -7, -3, 4,
	-2, 10, 5, -4, 9, 14, 8, 4, -1, 5,
	11, 12, 10, 7, 14, -1, 13, 5, -1, 13,
	5, 10, -1, 5, -1, 5,
}

var yyDef = [...]int{
	5, -2, 1, -2, 0, 6, 7, 2, 24, -2,
	0, 11, -2, 4, 25, 9, 0, 13, 14, 15,
	0, 0, 26, 0, 10, 16, 0, 20, 18, 0,
	22, 27, 17, 21, 19, 23,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int { _ = "STUB: not implemented"; return 0 }

func yyNewParser() yyParser { _ = "STUB: not implemented"; return *new(yyParser) }

const yyFlag = -1000

func yyTokname(c int) string { _ = "STUB: not implemented"; return "" }

func yyStatname(s int) string { _ = "STUB: not implemented"; return "" }

func yyErrorMessage(state, lookAhead int) string { _ = "STUB: not implemented"; return "" }

func yylex1(lex yyLexer, lval *yySymType) (char, token int) { _ = "STUB: not implemented"; return 0, 0 }

func yyParse(yylex yyLexer) int { _ = "STUB: not implemented"; return 0 }

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int { _ = "STUB: not implemented"; return 0 }

//line query_string.y:41

//line query_string.y:46

//line query_string.y:50

//line query_string.y:55

//line query_string.y:74

//line query_string.y:78

//line query_string.y:83

//line query_string.y:89

//line query_string.y:103

//line query_string.y:115

//line query_string.y:129

//line query_string.y:144

//line query_string.y:151

//line query_string.y:167

//line query_string.y:185

//line query_string.y:194

//line query_string.y:207

//line query_string.y:220

//line query_string.y:233

//line query_string.y:246

//line query_string.y:261

//line query_string.y:276

//line query_string.y:291

//line query_string.y:307

//line query_string.y:311

//line query_string.y:323

//line query_string.y:327

//line query_string.y:332

//line query_string.y:336
