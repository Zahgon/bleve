package scorch

import (
	"sync"

	"github.com/RoaringBitmap/roaring/v2"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type SegmentSnapshot struct {
	mmaped  uint32
	id      uint64
	segment segment.Segment
	deleted *roaring.Bitmap
	creator string
	stats   *fieldStats

	updatedFields map[string]*index.UpdateFieldInfo

	cachedMeta *cachedMeta

	cachedDocs *cachedDocs

	rootCountOnce sync.Once
	rootCount     uint64
}

func (s *SegmentSnapshot) Segment() segment.Segment {
	_ = "STUB: not implemented"
	return *new(segment.Segment)
}

func (s *SegmentSnapshot) Deleted() *roaring.Bitmap { _ = "STUB: not implemented"; return nil }

func (s *SegmentSnapshot) Id() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) FullSize() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) LiveSize() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) HasVector() bool { _ = "STUB: not implemented"; return false }

func (s *SegmentSnapshot) FileSize() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) LiveFileSize() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) Close() error { _ = "STUB: not implemented"; return nil }

func (s *SegmentSnapshot) VisitDocument(num uint64, visitor segment.StoredFieldValueVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SegmentSnapshot) DocID(num uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SegmentSnapshot) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) CountRoot() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) DocNumbers(docIDs []string) (*roaring.Bitmap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SegmentSnapshot) DocNumbersLive() *roaring.Bitmap { _ = "STUB: not implemented"; return nil }

func (s *SegmentSnapshot) Fields() []string { _ = "STUB: not implemented"; return nil }

func (s *SegmentSnapshot) Size() (rv int) { _ = "STUB: not implemented"; return 0 }

func (s *SegmentSnapshot) UpdateFieldsInfo(updatedFields map[string]*index.UpdateFieldInfo) {
	_ = "STUB: not implemented"
	return
}

type cachedFieldDocs struct {
	m       sync.Mutex
	readyCh chan struct{}
	err     error
	docs    map[uint64][]byte
	size    uint64
}

func (cfd *cachedFieldDocs) Size() int { _ = "STUB: not implemented"; return 0 }

func (cfd *cachedFieldDocs) prepareField(field string, ss *SegmentSnapshot) {
	_ = "STUB: not implemented"
	return
}

type cachedDocs struct {
	size  uint64
	m     sync.RWMutex
	cache map[string]*cachedFieldDocs
}

func (c *cachedDocs) prepareFields(wantedFields []string, ss *SegmentSnapshot) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cachedDocs) hasFields(fields []string) bool { _ = "STUB: not implemented"; return false }

func (c *cachedDocs) Size() int { _ = "STUB: not implemented"; return 0 }

func (c *cachedDocs) updateSizeLOCKED() { _ = "STUB: not implemented"; return }

func (c *cachedDocs) visitDoc(localDocNum uint64,
	fields []string, visitor index.DocValueVisitor) {
	_ = "STUB: not implemented"
	return
}

type cachedMeta struct {
	meta sync.Map
}

func newCachedMeta() *cachedMeta { _ = "STUB: not implemented"; return nil }

func (c *cachedMeta) store(field string, val interface{}) { _ = "STUB: not implemented"; return }

func (c *cachedMeta) load(field string) (rv interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *cachedMeta) contains(field string) bool { _ = "STUB: not implemented"; return false }

func (s *SegmentSnapshot) Ancestors(docNum uint64, prealloc []index.AncestorID) []index.AncestorID {
	_ = "STUB: not implemented"
	return nil
}
