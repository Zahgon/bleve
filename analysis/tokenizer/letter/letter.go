package letter

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "letter"

func TokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func init() {
	err := registry.RegisterTokenizer(Name, TokenizerConstructor)
	if err != nil {
		panic(err)
	}
}
