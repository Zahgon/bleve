package en

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const PluralStemmerName = "stemmer_en_plural"

type EnglishPluralStemmerFilter struct {
}

func NewEnglishPluralStemmerFilter() *EnglishPluralStemmerFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnglishPluralStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func EnglishPluralStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(PluralStemmerName, EnglishPluralStemmerFilterConstructor)
	if err != nil {
		panic(err)
	}
}

var oesExceptions = []string{"shoes", "canoes", "oboes"}

var chesExceptions = []string{
	"cliches",
	"avalanches",
	"mustaches",
	"moustaches",
	"quiches",
	"headaches",
	"heartaches",
	"porsches",
	"tranches",
	"caches",
}

func stem(word string) string { _ = "STUB: not implemented"; return "" }

func isException(word []rune, exceptions []string) bool { _ = "STUB: not implemented"; return false }
