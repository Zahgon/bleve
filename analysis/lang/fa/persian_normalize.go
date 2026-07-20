package fa

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const NormalizeName = "normalize_fa"

const (
	Yeh        = '\u064A'
	FarsiYeh   = '\u06CC'
	YehBarree  = '\u06D2'
	Keheh      = '\u06A9'
	Kaf        = '\u0643'
	HamzaAbove = '\u0654'
	HehYeh     = '\u06C0'
	HehGoal    = '\u06C1'
	Heh        = '\u0647'
)

type PersianNormalizeFilter struct {
}

func NewPersianNormalizeFilter() *PersianNormalizeFilter { _ = "STUB: not implemented"; return nil }

func (s *PersianNormalizeFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
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
