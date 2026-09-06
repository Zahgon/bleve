package ar

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const NormalizeName = "normalize_ar"

const (
	Alef           = '\u0627'
	AlefMadda      = '\u0622'
	AlefHamzaAbove = '\u0623'
	AlefHamzaBelow = '\u0625'
	Yeh            = '\u064A'
	DotlessYeh     = '\u0649'
	TehMarbuta     = '\u0629'
	Heh            = '\u0647'
	Tatweel        = '\u0640'
	Fathatan       = '\u064B'
	Dammatan       = '\u064C'
	Kasratan       = '\u064D'
	Fatha          = '\u064E'
	Damma          = '\u064F'
	Kasra          = '\u0650'
	Shadda         = '\u0651'
	Sukun          = '\u0652'
)

type ArabicNormalizeFilter struct {
}

func NewArabicNormalizeFilter() *ArabicNormalizeFilter { _ = "STUB: not implemented"; return nil }

func (s *ArabicNormalizeFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func normalize(input []byte) []byte { _ = "STUB: not implemented"; return nil }

func NormalizerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(NormalizeName, NormalizerFilterConstructor)
	if err != nil {
		panic(err)
	}
}
