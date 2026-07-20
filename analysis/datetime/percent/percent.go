package percent

import (
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "percentstyle"

var formatDelimiter byte = '%'

var formatSpecifierToLayout = map[byte]string{
	formatDelimiter: string(formatDelimiter),
	'a':             "Mon",
	'A':             "Monday",
	'd':             "02",
	'e':             "2",
	'b':             "Jan",
	'B':             "January",
	'm':             "01",
	'y':             "06",
	'Y':             "2006",
	'H':             "15",
	'I':             "03",
	'l':             "3",
	'p':             "PM",
	'P':             "pm",
	'M':             "04",
	'S':             "05",
	'f':             "999999",
	'Z':             "MST",

	'o': "1",
	'i': "4",
	's': "5",
	'N': "999999999",
}

var timezoneOptions = map[string]string{
	"z":   "Z0700",
	"z:M": "Z07:00",
	"z:S": "Z07:00:00",
	"zH":  "Z07",
	"zS":  "Z070000",
}

type DateTimeParser struct {
	layouts []string
}

func New(layouts []string) *DateTimeParser { _ = "STUB: not implemented"; return nil }

func checkTZOptions(formatString string, idx int) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

func parseFormatString(formatString string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *DateTimeParser) ParseDateTime(input string) (time.Time, string, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), "", nil
}

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
