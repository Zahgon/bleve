package iso

import (
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "isostyle"

var textLiteralDelimiter byte = '\''

var timeElementToLayout = map[byte]map[int]string{
	'M': {
		4: "January",
		3: "Jan",
		2: "01",
		1: "1",
	},
	'd': {
		2: "02",
		1: "2",
	},
	'a': {
		2: "pm",
		1: "PM",
	},
	'H': {
		2: "15",
		1: "15",
	},
	'm': {
		2: "04",
		1: "4",
	},
	's': {
		2: "05",
		1: "5",
	},

	'X': {
		5: "Z07:00:00",
		4: "Z070000",
		3: "Z07:00",
		2: "Z0700",
		1: "Z07",
	},
	'x': {
		5: "-07:00:00",
		4: "-070000",
		3: "-07:00",
		2: "-0700",
		1: "-07",
	},
}

type DateTimeParser struct {
	layouts []string
}

func New(layouts []string) *DateTimeParser { _ = "STUB: not implemented"; return nil }

func (p *DateTimeParser) ParseDateTime(input string) (time.Time, string, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), "", nil
}

func letterCounter(layout string, idx int) int { _ = "STUB: not implemented"; return 0 }

func invalidFormatError(character byte, count int) error { _ = "STUB: not implemented"; return nil }

func parseISOString(layout string) (string, error) { _ = "STUB: not implemented"; return "", nil }

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
