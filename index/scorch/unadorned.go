package scorch

import (
	"math"
	"reflect"

	"github.com/RoaringBitmap/roaring/v2"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

var reflectStaticSizeUnadornedPostingsIteratorBitmap int
var reflectStaticSizeUnadornedPostingsIterator1Hit int
var reflectStaticSizeUnadornedPosting int

func init() {
	var pib unadornedPostingsIteratorBitmap
	reflectStaticSizeUnadornedPostingsIteratorBitmap = int(reflect.TypeOf(pib).Size())
	var pi1h unadornedPostingsIterator1Hit
	reflectStaticSizeUnadornedPostingsIterator1Hit = int(reflect.TypeOf(pi1h).Size())
	var up UnadornedPosting
	reflectStaticSizeUnadornedPosting = int(reflect.TypeOf(up).Size())
}

type unadornedPostingsIteratorBitmap struct {
	actual   roaring.IntPeekable
	actualBM *roaring.Bitmap
	next     UnadornedPosting
}

func (i *unadornedPostingsIteratorBitmap) Next() (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (i *unadornedPostingsIteratorBitmap) Advance(docNum uint64) (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (i *unadornedPostingsIteratorBitmap) nextAtOrAfter(atOrAfter uint64) (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (i *unadornedPostingsIteratorBitmap) nextDocNumAtOrAfter(atOrAfter uint64) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (i *unadornedPostingsIteratorBitmap) Size() int { _ = "STUB: not implemented"; return 0 }

func (i *unadornedPostingsIteratorBitmap) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (i *unadornedPostingsIteratorBitmap) BytesWritten() uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (i *unadornedPostingsIteratorBitmap) ResetBytesRead(uint64) { _ = "STUB: not implemented"; return }

func (i *unadornedPostingsIteratorBitmap) ActualBitmap() *roaring.Bitmap {
	_ = "STUB: not implemented"
	return nil
}

func (i *unadornedPostingsIteratorBitmap) DocNum1Hit() (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (i *unadornedPostingsIteratorBitmap) ReplaceActual(actual *roaring.Bitmap) {
	_ = "STUB: not implemented"
	return
}

func (i *unadornedPostingsIteratorBitmap) ResetIterator() { _ = "STUB: not implemented"; return }

func newUnadornedPostingsIteratorFromBitmap(bm *roaring.Bitmap) segment.PostingsIterator {
	_ = "STUB: not implemented"
	return *new(segment.PostingsIterator)
}

const docNum1HitFinished = math.MaxUint64

type unadornedPostingsIterator1Hit struct {
	docNumOrig uint64
	docNum     uint64
	next       UnadornedPosting
}

func (i *unadornedPostingsIterator1Hit) Next() (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (i *unadornedPostingsIterator1Hit) Advance(docNum uint64) (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (i *unadornedPostingsIterator1Hit) nextAtOrAfter(atOrAfter uint64) (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (i *unadornedPostingsIterator1Hit) nextDocNumAtOrAfter(atOrAfter uint64) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (i *unadornedPostingsIterator1Hit) Size() int { _ = "STUB: not implemented"; return 0 }

func (i *unadornedPostingsIterator1Hit) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (i *unadornedPostingsIterator1Hit) BytesWritten() uint64 { _ = "STUB: not implemented"; return 0 }

func (i *unadornedPostingsIterator1Hit) ResetBytesRead(uint64) { _ = "STUB: not implemented"; return }

func (i *unadornedPostingsIterator1Hit) ResetIterator() { _ = "STUB: not implemented"; return }

func newUnadornedPostingsIteratorFrom1Hit(docNum1Hit uint64) segment.PostingsIterator {
	_ = "STUB: not implemented"
	return *new(segment.PostingsIterator)
}

type ResetablePostingsIterator interface {
	ResetIterator()
}

type UnadornedPosting struct {
	docNum uint64
}

func (p *UnadornedPosting) Number() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *UnadornedPosting) Frequency() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *UnadornedPosting) Norm() float64 { _ = "STUB: not implemented"; return 0 }

func (p *UnadornedPosting) Locations() []segment.Location { _ = "STUB: not implemented"; return nil }

func (p *UnadornedPosting) Size() int { _ = "STUB: not implemented"; return 0 }
