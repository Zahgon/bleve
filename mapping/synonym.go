package mapping

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

type SynonymSource struct {
	CollectionName string `json:"collection"`
	AnalyzerName   string `json:"analyzer"`
}

func NewSynonymSource(collection, analyzer string) *SynonymSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *SynonymSource) Collection() string { _ = "STUB: not implemented"; return "" }

func (s *SynonymSource) Analyzer() string { _ = "STUB: not implemented"; return "" }

func (s *SynonymSource) SetCollection(c string) { _ = "STUB: not implemented"; return }

func (s *SynonymSource) SetAnalyzer(a string) { _ = "STUB: not implemented"; return }

func SynonymSourceConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.SynonymSource, error) {
	_ = "STUB: not implemented"
	return *new(analysis.SynonymSource), nil
}

func init() {
	err := registry.RegisterSynonymSource(analysis.SynonymSourceType, SynonymSourceConstructor)
	if err != nil {
		panic(err)
	}
}
