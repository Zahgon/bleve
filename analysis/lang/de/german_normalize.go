package de

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const NormalizeName = "normalize_de"

const (
	N = 0
	V = 1
	U = 2
)

type GermanNormalizeFilter struct {
}

func NewGermanNormalizeFilter() *GermanNormalizeFilter { _ = "STUB: not implemented"; return nil }

func (s *GermanNormalizeFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
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
