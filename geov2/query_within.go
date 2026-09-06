package geov2

import (
	"github.com/blevesearch/bleve/v2/util"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type withinQuery struct {
	innerCells []uint64
	crossCells []uint64

	shape index.GeoJSON
	bBox  index.GeoJSON
}

func NewWithinQuery(shape index.GeoJSON) Query { _ = "STUB: not implemented"; return *new(Query) }

func (wq *withinQuery) Evaluate(geoData segment.GeoShapeV2Data) *util.Bitset {
	_ = "STUB: not implemented"
	return nil
}

func (wq *withinQuery) InnerCells() []uint64 { _ = "STUB: not implemented"; return nil }

func (wq *withinQuery) CrossCells() []uint64 { _ = "STUB: not implemented"; return nil }
