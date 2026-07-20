package single

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "single"

type SingleTokenTokenizer struct {
}

func NewSingleTokenTokenizer() *SingleTokenTokenizer { _ = "STUB: not implemented"; return nil }

func (t *SingleTokenTokenizer) Tokenize(input []byte) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func SingleTokenTokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func init() {
	err := registry.RegisterTokenizer(Name, SingleTokenTokenizerConstructor)
	if err != nil {
		panic(err)
	}
}
