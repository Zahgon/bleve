package scorer

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeTermQueryScorer int

func init() {
	var tqs TermQueryScorer
	reflectStaticSizeTermQueryScorer = int(reflect.TypeOf(tqs).Size())
}

type TermQueryScorer struct {
	queryTerm              string
	queryField             string
	queryBoost             float64
	docTerm                uint64
	docTotal               uint64
	avgDocLength           float64
	idf                    float64
	options                search.SearcherOptions
	idfExplanation         *search.Explanation
	includeScore           bool
	queryNorm              float64
	queryWeight            float64
	queryWeightExplanation *search.Explanation
}

func (s *TermQueryScorer) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *TermQueryScorer) computeIDF(avgDocLength float64, docTotal, docTerm uint64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func NewTermQueryScorer(queryTerm []byte, queryField string, queryBoost float64, docTotal,
	docTerm uint64, avgDocLength float64, options search.SearcherOptions) *TermQueryScorer {
	_ = "STUB: not implemented"
	return nil
}

func (s *TermQueryScorer) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (s *TermQueryScorer) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }

func (s *TermQueryScorer) docScore(tf, norm float64) (score float64, model string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (s *TermQueryScorer) scoreExplanation(tf float64, termMatch *index.TermFieldDoc) []*search.Explanation {
	_ = "STUB: not implemented"
	return nil
}

func (s *TermQueryScorer) Score(ctx *search.SearchContext, termMatch *index.TermFieldDoc) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}
