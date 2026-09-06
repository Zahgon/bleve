//go:build !vectors
// +build !vectors

package scorch

import (
	"github.com/blevesearch/bleve/v2/util"
	index "github.com/blevesearch/bleve_index_api"
)

func initTrainer(s *Scorch, config map[string]interface{}) *noopTrainer {
	_ = "STUB: not implemented"
	return nil
}

type noopTrainer struct {
}

func (t *noopTrainer) trainLoop() { _ = "STUB: not implemented"; return }

func (t *noopTrainer) train(batch *index.Batch) error { _ = "STUB: not implemented"; return nil }

func (t *noopTrainer) loadTrainedData(bucket *util.BoltBucketImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *noopTrainer) getInternal(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *noopTrainer) copyFileLOCKED(file string, d index.IndexDirectory) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *noopTrainer) updateBolt(snapshotsBucket *util.BoltBucketImpl, key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *noopTrainer) dropFileWriterIDs(ids map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *noopTrainer) fileWriterIDsInUse() (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
