package exception

import (
	"regexp"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "exception"

type ExceptionsTokenizer struct {
	exception *regexp.Regexp
	remaining analysis.Tokenizer
}

func NewExceptionsTokenizer(exception *regexp.Regexp, remaining analysis.Tokenizer) *ExceptionsTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (t *ExceptionsTokenizer) Tokenize(input []byte) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func ExceptionsTokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func init() {
	err := registry.RegisterTokenizer(Name, ExceptionsTokenizerConstructor)
	if err != nil {
		panic(err)
	}
}
