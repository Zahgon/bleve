//go:build vectors
// +build vectors

package scorch

import (
	"context"
	"encoding/json"

	"github.com/bits-and-blooms/bitset"
	index "github.com/blevesearch/bleve_index_api"
)

func (is *IndexSnapshot) VectorReader(ctx context.Context, vector []float32,
	field string, k int64, searchParams json.RawMessage,
	eligibleSelector index.EligibleDocumentSelector) (
	index.VectorReader, error) {
	_ = "STUB: not implemented"
	return *new(index.VectorReader), nil
}

type eligibleDocumentList struct {
	bs *bitset.BitSet
}

func (edl *eligibleDocumentList) Iterator() index.EligibleDocumentIterator {
	_ = "STUB: not implemented"
	return *new(index.EligibleDocumentIterator)
}

func (edl *eligibleDocumentList) Count() uint64 { _ = "STUB: not implemented"; return 0 }

var emptyEligibleDocumentList = &eligibleDocumentList{}

type eligibleDocumentIterator struct {
	bs      *bitset.BitSet
	current uint
}

func (it *eligibleDocumentIterator) Next() (id uint64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

var emptyEligibleIterator = &emptyEligibleDocumentIterator{}

type emptyEligibleDocumentIterator struct{}

func (it *emptyEligibleDocumentIterator) Next() (id uint64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type eligibleDocumentSelector struct {
	eligibleDocNums []*bitset.BitSet
	is              *IndexSnapshot
}

func (eds *eligibleDocumentSelector) SegmentEligibleDocuments(segmentID int) index.EligibleDocumentList {
	_ = "STUB: not implemented"
	return *new(index.EligibleDocumentList)
}

func (eds *eligibleDocumentSelector) AddEligibleDocumentMatch(id index.IndexInternalID) error {
	_ = "STUB: not implemented"
	return nil
}

func (is *IndexSnapshot) NewEligibleDocumentSelector() index.EligibleDocumentSelector {
	_ = "STUB: not implemented"
	return *new(index.EligibleDocumentSelector)
}
