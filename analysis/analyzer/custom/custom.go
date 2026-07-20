package custom

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "custom"

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer), nil
}

func init() {
	err := registry.RegisterAnalyzer(Name, AnalyzerConstructor)
	if err != nil {
		panic(err)
	}
}

func getCharFilters(charFilterNames []string, cache *registry.Cache) ([]analysis.CharFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTokenFilters(tokenFilterNames []string, cache *registry.Cache) ([]analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertInterfaceSliceToStringSlice(interfaceSlice []interface{}, objType string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
