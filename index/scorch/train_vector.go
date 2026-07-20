//go:build vectors
// +build vectors

package scorch

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/blevesearch/bleve/v2/util"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type trainRequest struct {
	finalSample          bool
	sampleSize           int
	ackCh                chan error
	sample               segment.Segment
	trainingParams       *index.TrainingParams
	removedFileWriterIDs map[string]struct{}
}

type vectorTrainer struct {
	trainingComplete atomic.Bool
	trainedSamples   uint64
	parent           *Scorch
	config           map[string]interface{}

	m sync.RWMutex

	trainedIndex *SegmentSnapshot
	trainCh      chan *trainRequest

	doneCh chan struct{}
}

const IndexTrainedWithFastMerge = "vector_index_fast_merge"

func initTrainer(s *Scorch, config map[string]interface{}) *vectorTrainer {
	_ = "STUB: not implemented"
	return nil
}

func moveFile(sourcePath, destPath string) error { _ = "STUB: not implemented"; return nil }

func (t *vectorTrainer) persistToBolt(trainReq *trainRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *vectorTrainer) trainLoop() { _ = "STUB: not implemented"; return }

func (t *vectorTrainer) recordTrainingStats(start time.Time) { _ = "STUB: not implemented"; return }

func ackRequest(req *trainRequest, err error) { _ = "STUB: not implemented"; return }

func (t *vectorTrainer) handleRequest(req *trainRequest, path string, config map[string]interface{}) (shutdown bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *vectorTrainer) publishTrainedIndex(req *trainRequest, prev *SegmentSnapshot, path string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *vectorTrainer) loadTrainedData(bucket *util.BoltBucketImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *vectorTrainer) train(batch *index.Batch) error { _ = "STUB: not implemented"; return nil }

func (t *vectorTrainer) submit(req *trainRequest) error { _ = "STUB: not implemented"; return nil }

func (t *vectorTrainer) getInternal(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *vectorTrainer) getTrainedIndex(field string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *vectorTrainer) copyFileLOCKED(file string, d index.IndexDirectory) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *vectorTrainer) updateBolt(snapshotsBucket *util.BoltBucketImpl, key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *vectorTrainer) dropFileWriterIDs(ids map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *vectorTrainer) removeFileWriterIDs(ids map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *vectorTrainer) fileWriterIDsInUse() (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
