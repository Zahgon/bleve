package fusion

import (
	"github.com/blevesearch/bleve/v2/search"
)

type FusionResult struct {
	Hits     search.DocumentMatchCollection
	Total    uint64
	MaxScore float64
}
