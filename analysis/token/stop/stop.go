package stop

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "stop_tokens"

type StopTokensFilter struct {
	stopTokens analysis.TokenMap
}

func NewStopTokensFilter(stopTokens analysis.TokenMap) *StopTokensFilter {
	_ = "STUB: not implemented"
	return nil
}

func (f *StopTokensFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func StopTokensFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, StopTokensFilterConstructor)
	if err != nil {
		panic(err)
	}
}
