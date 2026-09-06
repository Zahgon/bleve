package scorch

import (
	"reflect"

	"github.com/RoaringBitmap/roaring/v2"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeIndexSnapshotDocIDReader int

func init() {
	var isdr IndexSnapshotDocIDReader
	reflectStaticSizeIndexSnapshotDocIDReader = int(reflect.TypeOf(isdr).Size())
}

type IndexSnapshotDocIDReader struct {
	snapshot      *IndexSnapshot
	iterators     []roaring.IntIterable
	segmentOffset int
}

func (i *IndexSnapshotDocIDReader) Size() int { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotDocIDReader) Next() (index.IndexInternalID, error) {
	_ = "STUB: not implemented"
	return *new(index.IndexInternalID), nil
}

func (i *IndexSnapshotDocIDReader) Advance(ID index.IndexInternalID) (index.IndexInternalID, error) {
	_ = "STUB: not implemented"
	return *new(index.IndexInternalID), nil
}

func (i *IndexSnapshotDocIDReader) Close() error { _ = "STUB: not implemented"; return nil }
