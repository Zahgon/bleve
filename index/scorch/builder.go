package scorch

import (
	"sync"

	index "github.com/blevesearch/bleve_index_api"
)

const DefaultBuilderBatchSize = 1000
const DefaultBuilderMergeMax = 10

type Builder struct {
	m         sync.Mutex
	segCount  uint64
	path      string
	buildPath string
	segPaths  []string
	batchSize int
	mergeMax  int
	batch     *index.Batch
	internal  map[string][]byte
	segPlugin SegmentPlugin
}

func NewBuilder(config map[string]interface{}) (*Builder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Builder) parseConfig(config map[string]interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (o *Builder) Index(doc index.Document) error { _ = "STUB: not implemented"; return nil }

func (o *Builder) maybeFlushBatchLOCKED(moreThan int) error { _ = "STUB: not implemented"; return nil }

func (o *Builder) executeBatchLOCKED(batch *index.Batch) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (o *Builder) doMerge() error { _ = "STUB: not implemented"; return nil }

func (o *Builder) Close() error { _ = "STUB: not implemented"; return nil }
