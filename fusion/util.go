package fusion

import (
	"github.com/blevesearch/bleve/v2/search"
)

func sortDocMatchesByScore(hits search.DocumentMatchCollection) { _ = "STUB: not implemented"; return }

func scoreBreakdownForQuery(hit *search.DocumentMatch, idx int) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func sortDocMatchesByBreakdown(hits search.DocumentMatchCollection, queryIdx int) {
	_ = "STUB: not implemented"
	return
}

func getFusionExplAt(hit *search.DocumentMatch, i int, value float64, message string) *search.Explanation {
	_ = "STUB: not implemented"
	return nil
}

func finalizeFusionExpl(hit *search.DocumentMatch, explChildren []*search.Explanation) {
	_ = "STUB: not implemented"
	return
}
