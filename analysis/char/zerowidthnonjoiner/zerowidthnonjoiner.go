package zerowidthnonjoiner

import (
	"regexp"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "zero_width_spaces"

var zeroWidthNonJoinerRegexp = regexp.MustCompile(`\x{200C}`)

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
