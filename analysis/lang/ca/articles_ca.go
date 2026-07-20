package ca

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const ArticlesName = "articles_ca"

var CatalanArticles = []byte(`
d
l
m
n
s
t
`)

func ArticlesTokenMapConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenMap, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenMap), nil
}

func init() {
	err := registry.RegisterTokenMap(ArticlesName, ArticlesTokenMapConstructor)
	if err != nil {
		panic(err)
	}
}
