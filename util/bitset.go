package util

import (
	"github.com/RoaringBitmap/roaring/v2"
)

type Bitset struct {
	data    []uint64
	numBits int
	exclude *roaring.Bitmap
}

func NewBitset(maxVal int, exclude *roaring.Bitmap) *Bitset { _ = "STUB: not implemented"; return nil }

func (b *Bitset) Add(val int) { _ = "STUB: not implemented"; return }

func (b *Bitset) Remove(val int) { _ = "STUB: not implemented"; return }

func (b *Bitset) Contains(val int) bool { _ = "STUB: not implemented"; return false }

func (b *Bitset) Invert() { _ = "STUB: not implemented"; return }

func (b *Bitset) Iterate(f func(int)) { _ = "STUB: not implemented"; return }

func (b *Bitset) Count() int { _ = "STUB: not implemented"; return 0 }

func (b *Bitset) Clear() { _ = "STUB: not implemented"; return }
