package regexp

import (
	"regexp"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "regexp"

type CharFilter struct {
	r           *regexp.Regexp
	replacement []byte
}

func New(r *regexp.Regexp, replacement []byte) *CharFilter { _ = "STUB: not implemented"; return nil }

func (s *CharFilter) Filter(input []byte) []byte { _ = "STUB: not implemented"; return nil }

func CharFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.CharFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.CharFilter), nil
}

func init() {
	err := registry.RegisterCharFilter(Name, CharFilterConstructor)
	if err != nil {
		panic(err)
	}
}
