package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func RegisterAnalyzer(name string, constructor AnalyzerConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type AnalyzerConstructor func(config map[string]interface{}, cache *Cache) (analysis.Analyzer, error)
type AnalyzerRegistry map[string]AnalyzerConstructor

type AnalyzerCache struct {
	*ConcurrentCache
}

func NewAnalyzerCache() *AnalyzerCache { _ = "STUB: not implemented"; return nil }

func AnalyzerBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *AnalyzerCache) AnalyzerNamed(name string, cache *Cache) (analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer), nil
}

func (c *AnalyzerCache) DefineAnalyzer(name string, typ string, config map[string]interface{}, cache *Cache) (analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer), nil
}

func AnalyzerTypesAndInstances() ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }
