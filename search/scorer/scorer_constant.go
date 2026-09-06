package scorer

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeConstantScorer int

func init() {
	var cs ConstantScorer
	reflectStaticSizeConstantScorer = int(reflect.TypeOf(cs).Size())
}

type ConstantScorer struct {
	constant               float64
	boost                  float64
	options                search.SearcherOptions
	queryNorm              float64
	queryWeight            float64
	queryWeightExplanation *search.Explanation
	includeScore           bool
}

func (s *ConstantScorer) Size() int { _ = "STUB: not implemented"; return 0 }

func NewConstantScorer(constant float64, boost float64, options search.SearcherOptions) *ConstantScorer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ConstantScorer) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *ConstantScorer) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *ConstantScorer) Score(ctx *search.SearchContext, id index.IndexInternalID) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}
