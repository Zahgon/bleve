package microseconds

import (
	"math"
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "unix_micro"

type DateTimeParser struct {
}

var minBound int64 = math.MinInt64 / 1000
var maxBound int64 = math.MaxInt64 / 1000

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
