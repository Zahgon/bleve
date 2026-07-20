package regexp

import (
	"regexp"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "regexp"

var IdeographRegexp = regexp.MustCompile(`\p{Han}|\p{Hangul}|\p{Hiragana}|\p{Katakana}`)

type RegexpTokenizer struct {
	r *regexp.Regexp
}

func NewRegexpTokenizer(r *regexp.Regexp) *RegexpTokenizer { _ = "STUB: not implemented"; return nil }

func (rt *RegexpTokenizer) Tokenize(input []byte) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func RegexpTokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func init() {
	err := registry.RegisterTokenizer(Name, RegexpTokenizerConstructor)
	if err != nil {
		panic(err)
	}
}

func detectTokenType(termBytes []byte) analysis.TokenType {
	_ = "STUB: not implemented"
	return *new(analysis.TokenType)
}
