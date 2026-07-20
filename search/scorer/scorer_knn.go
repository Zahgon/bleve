//go:build vectors
// +build vectors

package scorer

import (
	"math"
	"reflect"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeKNNQueryScorer int

func init() {
	var sqs KNNQueryScorer
	reflectStaticSizeKNNQueryScorer = int(reflect.TypeOf(sqs).Size())
}

type KNNQueryScorer struct {
	queryVector            []float32
	queryField             string
	queryWeight            float64
	queryBoost             float64
	queryNorm              float64
	options                search.SearcherOptions
	similarityMetric       string
	queryWeightExplanation *search.Explanation
}

func (s *KNNQueryScorer) Size() int { _ = "STUB: not implemented"; return 0 }

func NewKNNQueryScorer(queryVector []float32, queryField string, queryBoost float64,
	options search.SearcherOptions,
	similarityMetric string) *KNNQueryScorer {
	_ = "STUB: not implemented"
	return nil
}

const maxKNNScore = math.MaxFloat32

func (sqs *KNNQueryScorer) Score(ctx *search.SearchContext,
	knnMatch *index.VectorDoc) *search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func (sqs *KNNQueryScorer) Weight() float64 { _ = "STUB: not implemented"; return 0 }

func (sqs *KNNQueryScorer) SetQueryNorm(qnorm float64) { _ = "STUB: not implemented"; return }
