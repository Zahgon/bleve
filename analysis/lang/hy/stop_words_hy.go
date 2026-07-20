package hy

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const StopName = "stop_hy"

var ArmenianStopWords = []byte(`# example set of Armenian stopwords.
այդ
այլ
այն
այս
դու
դուք
եմ
են
ենք
ես
եք
է
էի
էին
էինք
էիր
էիք
էր
ըստ
թ
ի
ին
իսկ
իր
կամ
համար
հետ
հետո
մենք
մեջ
մի
ն
նա
նաև
նրա
նրանք
որ
որը
որոնք
որպես
ու
ում
պիտի
վրա
և
`)

func TokenMapConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenMap, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenMap), nil
}

func init() {
	err := registry.RegisterTokenMap(StopName, TokenMapConstructor)
	if err != nil {
		panic(err)
	}
}
