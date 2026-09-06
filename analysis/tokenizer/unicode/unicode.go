package unicode

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "unicode"

type UnicodeTokenizer struct {
}

func NewUnicodeTokenizer() *UnicodeTokenizer { _ = "STUB: not implemented"; return nil }

func (rt *UnicodeTokenizer) Tokenize(input []byte) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func UnicodeTokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func init() {
	err := registry.RegisterTokenizer(Name, UnicodeTokenizerConstructor)
	if err != nil {
		panic(err)
	}
}

func convertType(segmentWordType int) analysis.TokenType {
	_ = "STUB: not implemented"
	return *new(analysis.TokenType)
}
