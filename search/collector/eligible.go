//go:build vectors
// +build vectors

package collector

import (
	"context"
	"time"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type EligibleCollector struct {
	size             int
	total            uint64
	took             time.Duration
	eligibleSelector index.EligibleDocumentSelector
}

func NewEligibleCollector(size int) *EligibleCollector { _ = "STUB: not implemented"; return nil }

func newEligibleCollector(size int) *EligibleCollector { _ = "STUB: not implemented"; return nil }

func makeEligibleDocumentMatchHandler(ctx *search.SearchContext, reader index.IndexReader) (search.DocumentMatchHandler, error) {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchHandler), nil
}

func (ec *EligibleCollector) Collect(ctx context.Context, searcher search.Searcher, reader index.IndexReader) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *EligibleCollector) Results() search.DocumentMatchCollection {
	_ = "STUB: not implemented"
	return *new(search.DocumentMatchCollection)
}

func (ec *EligibleCollector) EligibleSelector() index.EligibleDocumentSelector {
	_ = "STUB: not implemented"
	return *new(index.EligibleDocumentSelector)
}

func (ec *EligibleCollector) Total() uint64 { _ = "STUB: not implemented"; return 0 }

func (ec *EligibleCollector) MaxScore() float64 { _ = "STUB: not implemented"; return 0 }

func (ec *EligibleCollector) Took() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ec *EligibleCollector) SetFacetsBuilder(facetsBuilder *search.FacetsBuilder) {
	_ = "STUB: not implemented"
	return
}

func (ec *EligibleCollector) FacetResults() search.FacetResults {
	_ = "STUB: not implemented"
	return *new(search.FacetResults)
}
