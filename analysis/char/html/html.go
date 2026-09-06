package html

import (
	"regexp"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "html"

var htmlCharFilterRegexp = regexp.MustCompile(`</?[!\w]+((\s+\w+(\s*=\s*(?:".*?"|'.*?'|[^'">\s]+))?)+\s*|\s*)/?>`)

type CharFilter struct {
	r           *regexp.Regexp
	replacement []byte
}

func New() *CharFilter { _ = "STUB: not implemented"; return nil }

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
