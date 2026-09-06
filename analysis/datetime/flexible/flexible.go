package flexible

import (
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "flexiblego"

type DateTimeParser struct {
	layouts []string
}

func New(layouts []string) *DateTimeParser { _ = "STUB: not implemented"; return nil }

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
