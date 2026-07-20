package nl

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const AnalyzerName = "nl"

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer), nil
}

func init() {
	err := registry.RegisterAnalyzer(AnalyzerName, AnalyzerConstructor)
	if err != nil {
		panic(err)
	}
}
