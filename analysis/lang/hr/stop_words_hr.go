package hr

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const StopName = "stop_hr"

var CroatianStopWords = []byte(`biti
jesam
budem
sam
jesi
budeš
si
jesmo
budemo
smo
jeste
budete
ste
jesu
budu
su
bih
bijah
bjeh
bijaše
bi
bje
bješe
bijasmo
bismo
bjesmo
bijaste
biste
bjeste
bijahu
biste
bjeste
bijahu
bi
biše
bjehu
bješe
bio
bili
budimo
budite
bila
bilo
bile
ću
ćeš
će
ćemo
ćete
želim
želiš
želi
želimo
želite
žele
moram
moraš
mora
moramo
morate
moraju
trebam
trebaš
treba
trebamo
trebate
trebaju
mogu
možeš
može
možemo
možete
za
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
