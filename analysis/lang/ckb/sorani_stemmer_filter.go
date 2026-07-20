package ckb

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const StemmerName = "stemmer_ckb"

type SoraniStemmerFilter struct {
}

func NewSoraniStemmerFilter() *SoraniStemmerFilter { _ = "STUB: not implemented"; return nil }

func (s *SoraniStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func stem(input []byte) []byte { _ = "STUB: not implemented"; return nil }

func truncateRunes(input []byte, num int) []byte { _ = "STUB: not implemented"; return nil }

func buildTermFromRunes(runes []rune) []byte { _ = "STUB: not implemented"; return nil }

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
