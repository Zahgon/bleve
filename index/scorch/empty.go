package scorch

import (
	"github.com/RoaringBitmap/roaring/v2"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type emptyPostingsIterator struct{}

func (e *emptyPostingsIterator) Next() (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (e *emptyPostingsIterator) Advance(uint64) (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (e *emptyPostingsIterator) Size() int { _ = "STUB: not implemented"; return 0 }

func (e *emptyPostingsIterator) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (e *emptyPostingsIterator) ResetBytesRead(uint64) { _ = "STUB: not implemented"; return }

func (e *emptyPostingsIterator) BytesWritten() uint64 { _ = "STUB: not implemented"; return 0 }

func (e *emptyPostingsIterator) ActualBitmap() *roaring.Bitmap {
	_ = "STUB: not implemented"
	return nil
}

func (e *emptyPostingsIterator) DocNum1Hit() (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (e *emptyPostingsIterator) ReplaceActual(*roaring.Bitmap) { _ = "STUB: not implemented"; return }

var anEmptyPostingsIterator = &emptyPostingsIterator{}
