package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func RegisterDateTimeParser(name string, constructor DateTimeParserConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type DateTimeParserConstructor func(config map[string]interface{}, cache *Cache) (analysis.DateTimeParser, error)
type DateTimeParserRegistry map[string]DateTimeParserConstructor

type DateTimeParserCache struct {
	*ConcurrentCache
}

func NewDateTimeParserCache() *DateTimeParserCache { _ = "STUB: not implemented"; return nil }

func DateTimeParserBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *DateTimeParserCache) DateTimeParserNamed(name string, cache *Cache) (analysis.DateTimeParser, error) {
	_ = "STUB: not implemented"
	return *new(analysis.DateTimeParser), nil
}

func (c *DateTimeParserCache) DefineDateTimeParser(name string, typ string, config map[string]interface{}, cache *Cache) (analysis.DateTimeParser, error) {
	_ = "STUB: not implemented"
	return *new(analysis.DateTimeParser), nil
}

func DateTimeParserTypesAndInstances() ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
