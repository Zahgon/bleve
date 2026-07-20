package en

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const PossessiveName = "possessive_en"

const rightSingleQuotationMark = '’'
const apostrophe = '\''
const fullWidthApostrophe = '＇'

const apostropheChars = rightSingleQuotationMark + apostrophe + fullWidthApostrophe

type PossessiveFilter struct {
}

func NewPossessiveFilter() *PossessiveFilter { _ = "STUB: not implemented"; return nil }

func (s *PossessiveFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func PossessiveFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(PossessiveName, PossessiveFilterConstructor)
	if err != nil {
		panic(err)
	}
}
