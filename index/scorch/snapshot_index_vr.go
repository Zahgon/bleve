//go:build vectors
// +build vectors

package scorch

import (
	"context"
	"encoding/json"
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
	segment_api "github.com/blevesearch/scorch_segment_api/v2"
)

const VectorSearchSupportedSegmentVersion = 16

var reflectStaticSizeIndexSnapshotVectorReader int

func init() {
	var istfr IndexSnapshotVectorReader
	reflectStaticSizeIndexSnapshotVectorReader = int(reflect.TypeOf(istfr).Size())
}

type IndexSnapshotVectorReader struct {
	vector        []float32
	field         string
	k             int64
	snapshot      *IndexSnapshot
	postings      []segment_api.VecPostingsList
	iterators     []segment_api.VecPostingsIterator
	segmentOffset int
	currPosting   segment_api.VecPosting
	currID        index.IndexInternalID
	ctx           context.Context

	searchParams     json.RawMessage
	eligibleSelector index.EligibleDocumentSelector
}

func (i *IndexSnapshotVectorReader) Size() int { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotVectorReader) Next(preAlloced *index.VectorDoc) (
	*index.VectorDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshotVectorReader) Advance(ID index.IndexInternalID,
	preAlloced *index.VectorDoc) (*index.VectorDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshotVectorReader) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotVectorReader) Close() error { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshot) CentroidCardinalities(field string, limit int, descending bool) (
	[]index.CentroidCardinality, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
