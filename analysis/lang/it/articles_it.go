package it

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const ArticlesName = "articles_it"

var ItalianArticles = []byte(`
c
l
all
dall
dell
nell
sull
coll
pell
gl
agl
dagl
degl
negl
sugl
un
m
t
s
v
d
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
