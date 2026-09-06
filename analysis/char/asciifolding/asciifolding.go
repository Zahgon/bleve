package asciifolding

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "asciifolding"

type AsciiFoldingFilter struct{}

func New() *AsciiFoldingFilter { _ = "STUB: not implemented"; return nil }

func (s *AsciiFoldingFilter) Filter(input []byte) []byte { _ = "STUB: not implemented"; return nil }

func AsciiFoldingFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.CharFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.CharFilter), nil
}

func init() {
	err := registry.RegisterCharFilter(Name, AsciiFoldingFilterConstructor)
	if err != nil {
		panic(err)
	}
}

func foldToASCII(input []rune, inputPos int, output []rune, outputPos int, length int) []rune {
	_ = "STUB: not implemented"
	return nil
}
