package whitespace

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "whitespace"

func TokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func notSpace(r rune) bool { _ = "STUB: not implemented"; return false }

func init() {
	err := registry.RegisterTokenizer(Name, TokenizerConstructor)
	if err != nil {
		panic(err)
	}
}
