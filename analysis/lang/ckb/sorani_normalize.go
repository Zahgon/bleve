package ckb

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const NormalizeName = "normalize_ckb"

const (
	Yeh        = '\u064A'
	DotlessYeh = '\u0649'
	FarsiYeh   = '\u06CC'

	Kaf   = '\u0643'
	Keheh = '\u06A9'

	Heh            = '\u0647'
	Ae             = '\u06D5'
	Zwnj           = '\u200C'
	HehDoachashmee = '\u06BE'
	TehMarbuta     = '\u0629'

	Reh       = '\u0631'
	Rreh      = '\u0695'
	RrehAbove = '\u0692'

	Tatweel  = '\u0640'
	Fathatan = '\u064B'
	Dammatan = '\u064C'
	Kasratan = '\u064D'
	Fatha    = '\u064E'
	Damma    = '\u064F'
	Kasra    = '\u0650'
	Shadda   = '\u0651'
	Sukun    = '\u0652'
)

type SoraniNormalizeFilter struct {
}

func NewSoraniNormalizeFilter() *SoraniNormalizeFilter { _ = "STUB: not implemented"; return nil }

func (s *SoraniNormalizeFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
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
