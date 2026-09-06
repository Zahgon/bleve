package scorch

import (
	"github.com/RoaringBitmap/roaring/v2"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

const introducer = "introducer"

type segmentIntroduction struct {
	id        uint64
	data      segment.Segment
	obsoletes map[uint64]*roaring.Bitmap
	ids       []string
	internal  map[string][]byte

	applied           chan error
	persisted         chan error
	persistedCallback index.BatchCallback
}

type persistIntroduction struct {
	persisted map[uint64]segment.Segment
	applied   notificationChan
}

type epochWatcher struct {
	epoch    uint64
	notifyCh notificationChan
}

func (s *Scorch) introducerLoop() { _ = "STUB: not implemented"; return }

func (s *Scorch) introduceSegment(next *segmentIntroduction) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scorch) introducePersist(persist *persistIntroduction) { _ = "STUB: not implemented"; return }

func (s *Scorch) introduceMerge(nextMerge *segmentMerge) { _ = "STUB: not implemented"; return }

func isMemorySegment(s *SegmentSnapshot) bool { _ = "STUB: not implemented"; return false }
