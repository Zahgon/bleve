package fusion

import (
	"github.com/blevesearch/bleve/v2/search"
)

func formatRRFMessage(weight float64, rank int, rankConstant int) string {
	_ = "STUB: not implemented"
	return ""
}

func ReciprocalRankFusion(hits search.DocumentMatchCollection, weights []float64, rankConstant int, windowSize int, numKNNQueries int, explain bool) *FusionResult {
	_ = "STUB: not implemented"
	return nil
}
