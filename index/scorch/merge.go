package scorch

import (
	"context"

	"github.com/RoaringBitmap/roaring/v2"
	"github.com/blevesearch/bleve/v2/index/scorch/mergeplan"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

const merger = "merger"

const mergeDoneKey = "mergeDone"

type mergeDoneChan chan error

const mergePlanFuncKey = "mergePlanFunc"

type mergePlanFunc func(*IndexSnapshot) (*mergeplan.MergePlan, error)

func (s *Scorch) mergerLoop() { _ = "STUB: not implemented"; return }

type mergerCtrl struct {
	ctx     context.Context
	options *mergeplan.MergePlanOptions
	doneCh  chan struct{}
}

func (s *Scorch) ForceMerge(ctx context.Context,
	mo *mergeplan.MergePlanOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scorch) parseMergePlannerOptions(po *persisterOptions) (*mergeplan.MergePlanOptions,
	error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type closeChWrapper struct {
	ch1      chan struct{}
	ctx      context.Context
	closeCh  chan struct{}
	cancelCh chan struct{}
}

func newCloseChWrapper(ch1 chan struct{},
	ctx context.Context) *closeChWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (w *closeChWrapper) close() { _ = "STUB: not implemented"; return }

func (w *closeChWrapper) listen() { _ = "STUB: not implemented"; return }

type mergeBatch struct {
	snapshots []*SegmentSnapshot
	filenames []string
	segments  []segment.Segment
	drops     []*roaring.Bitmap

	new         segment.Segment
	newID       uint64
	newDocNums  [][]uint64
	newFilename string
	newTime     uint64
}

func (s *Scorch) planMergeAtSnapshot(ctrlMsg *mergerCtrl, ourSnapshot *IndexSnapshot) error {
	_ = "STUB: not implemented"
	return nil
}

type mergeTaskIntroStatus struct {
	indexSnapshot *IndexSnapshot
	skipped       []bool
}

type mergedSegmentHistory struct {
	batchID      int
	oldNewDocIDs []uint64
	oldSegment   *SegmentSnapshot
}

type segmentMerge struct {
	newSegmentIDs    []uint64
	newSegments      []segment.Segment
	mergedSegHistory map[uint64]*mergedSegmentHistory
	notifyCh         chan *mergeTaskIntroStatus
	mmaped           uint32
	fileMerge        bool
}

func cumulateBytesRead(sbs []segment.Segment) uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Scorch) mergeAndPersistInMemorySegments(flushes []*flushable, po *persisterOptions) (*IndexSnapshot, map[uint64]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *Scorch) ReportBytesWritten(bytesWritten uint64) { _ = "STUB: not implemented"; return }
