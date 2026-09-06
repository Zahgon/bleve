package optional

import (
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "dateTimeOptional"

const rfc3339NoTimezone = "2006-01-02T15:04:05"
const rfc3339NoTimezoneNoT = "2006-01-02 15:04:05"
const rfc3339Offset = "2006-01-02 15:04:05 -0700"
const rfc3339NoTime = "2006-01-02"

var layouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	rfc3339NoTimezone,
	rfc3339NoTimezoneNoT,
	rfc3339Offset,
	rfc3339NoTime,
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
