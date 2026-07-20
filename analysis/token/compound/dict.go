package compound

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "dict_compound"

const defaultMinWordSize = 5
const defaultMinSubWordSize = 2
const defaultMaxSubWordSize = 15
const defaultOnlyLongestMatch = false

type DictionaryCompoundFilter struct {
	dict             analysis.TokenMap
	minWordSize      int
	minSubWordSize   int
	maxSubWordSize   int
	onlyLongestMatch bool
}

func NewDictionaryCompoundFilter(dict analysis.TokenMap, minWordSize, minSubWordSize, maxSubWordSize int, onlyLongestMatch bool) *DictionaryCompoundFilter {
	_ = "STUB: not implemented"
	return nil
}

func (f *DictionaryCompoundFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func (f *DictionaryCompoundFilter) decompose(token *analysis.Token) []*analysis.Token {
	_ = "STUB: not implemented"
	return nil
}

func DictionaryCompoundFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, DictionaryCompoundFilterConstructor)
	if err != nil {
		panic(err)
	}
}
