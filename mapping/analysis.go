package mapping

type customAnalysis struct {
	CharFilters     map[string]map[string]interface{} `json:"char_filters,omitempty"`
	Tokenizers      map[string]map[string]interface{} `json:"tokenizers,omitempty"`
	TokenMaps       map[string]map[string]interface{} `json:"token_maps,omitempty"`
	TokenFilters    map[string]map[string]interface{} `json:"token_filters,omitempty"`
	Analyzers       map[string]map[string]interface{} `json:"analyzers,omitempty"`
	DateTimeParsers map[string]map[string]interface{} `json:"date_time_parsers,omitempty"`
	SynonymSources  map[string]map[string]interface{} `json:"synonym_sources,omitempty"`
}

func (c *customAnalysis) registerAll(i *IndexMappingImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func newCustomAnalysis() *customAnalysis { _ = "STUB: not implemented"; return nil }
