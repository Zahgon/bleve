//go:build vectors
// +build vectors

package scorch

import (
	"context"

	index "github.com/blevesearch/bleve_index_api"
)

type OptimizeVR struct {
	ctx       context.Context
	snapshot  *IndexSnapshot
	totalCost uint64

	vrs map[string][]*IndexSnapshotVectorReader
}

func (o *OptimizeVR) invokeSearcherEndCallback() { _ = "STUB: not implemented"; return }

func (o *OptimizeVR) search(segID int) error { _ = "STUB: not implemented"; return nil }

func (o *OptimizeVR) Finish() error { _ = "STUB: not implemented"; return nil }

func (s *IndexSnapshotVectorReader) VectorOptimize(ctx context.Context,
	octx index.VectorOptimizableContext,
) (index.VectorOptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.VectorOptimizableContext), nil
}
