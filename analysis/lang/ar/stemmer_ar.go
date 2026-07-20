package ar

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const StemmerName = "stemmer_ar"

var prefixes = [][]rune{
	[]rune("ال"),
	[]rune("وال"),
	[]rune("بال"),
	[]rune("كال"),
	[]rune("فال"),
	[]rune("لل"),
	[]rune("و"),
}
var suffixes = [][]rune{
	[]rune("ها"),
	[]rune("ان"),
	[]rune("ات"),
	[]rune("ون"),
	[]rune("ين"),
	[]rune("يه"),
	[]rune("ية"),
	[]rune("ه"),
	[]rune("ة"),
	[]rune("ي"),
}

type ArabicStemmerFilter struct{}

func NewArabicStemmerFilter() *ArabicStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *ArabicStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func canStemPrefix(input, prefix []rune) bool { _ = "STUB: not implemented"; return false }

func canStemSuffix(input, suffix []rune) bool { _ = "STUB: not implemented"; return false }

func stem(input []byte) []byte { _ = "STUB: not implemented"; return nil }

func StemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(StemmerName, StemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
