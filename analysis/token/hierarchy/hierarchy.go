package hierarchy

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "hierarchy"

type HierarchyFilter struct {
	maxLevels  int
	delimiter  []byte
	splitInput bool
}

func NewHierarchyFilter(delimiter []byte, maxLevels int, splitInput bool) *HierarchyFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *HierarchyFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func (s *HierarchyFilter) buildToken(tokenStream analysis.TokenStream, soFar [][]byte, part []byte) (
	[][]byte, analysis.TokenStream) {
	_ = "STUB: not implemented"
	return nil, *new(analysis.TokenStream)
}

func HierarchyFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, HierarchyFilterConstructor)
	if err != nil {
		panic(err)
	}
}
