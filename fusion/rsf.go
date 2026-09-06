package fusion

import (
	"github.com/blevesearch/bleve/v2/search"
)

func formatRSFMessage(weight float64, normalizedScore float64, minScore float64, maxScore float64) string {
	_ = "STUB: not implemented"
	return ""
}

func RelativeScoreFusion(hits search.DocumentMatchCollection, weights []float64, windowSize int, numKNNQueries int, explain bool) *FusionResult {
	_ = "STUB: not implemented"
	return nil
}
