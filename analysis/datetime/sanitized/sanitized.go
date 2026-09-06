package sanitized

import (
	"regexp"
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "sanitizedgo"

var validMagicNumbers = map[string]struct{}{
	"2006":    {},
	"06":      {},
	"01":      {},
	"1":       {},
	"_1":      {},
	"January": {},
	"Jan":     {},
	"02":      {},
	"2":       {},
	"_2":      {},
	"__2":     {},
	"002":     {},
	"Monday":  {},
	"Mon":     {},
	"15":      {},
	"3":       {},
	"03":      {},
	"4":       {},
	"04":      {},
	"5":       {},
	"05":      {},
	"0700":    {},
	"070000":  {},
	"07":      {},
	"00":      {},
	"":        {},
}

var layoutSplitRegex = regexp.MustCompile("[\\+\\-= :T,Z\\.<>;\\?!`~@#$%\\^&\\*|'\"\\(\\){}\\[\\]/\\\\]")

var layoutStripRegex = regexp.MustCompile(`PM|pm|\.9+|\.0+|MST`)

type DateTimeParser struct {
	layouts []string
}

func New(layouts []string) *DateTimeParser { _ = "STUB: not implemented"; return nil }

func (p *DateTimeParser) ParseDateTime(input string) (time.Time, string, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), "", nil
}

func validateLayout(layout string) bool { _ = "STUB: not implemented"; return false }

func DateTimeParserConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.DateTimeParser, error) {
	_ = "STUB: not implemented"
	return *new(analysis.DateTimeParser), nil
}

func init() {
	err := registry.RegisterDateTimeParser(Name, DateTimeParserConstructor)
	if err != nil {
		panic(err)
	}
}
