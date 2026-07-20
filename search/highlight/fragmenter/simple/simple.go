package simple

import (
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search/highlight"
)

const Name = "simple"

const defaultFragmentSize = 200

type Fragmenter struct {
	fragmentSize int
}

func NewFragmenter(fragmentSize int) *Fragmenter { _ = "STUB: not implemented"; return nil }

func (s *Fragmenter) Fragment(orig []byte, ot highlight.TermLocations) []*highlight.Fragment {
	_ = "STUB: not implemented"
	return nil
}

func Constructor(config map[string]interface{}, cache *registry.Cache) (highlight.Fragmenter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Fragmenter), nil
}

func init() {
	err := registry.RegisterFragmenter(Name, Constructor)
	if err != nil {
		panic(err)
	}
}
