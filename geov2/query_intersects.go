package geov2

import (
	"github.com/blevesearch/bleve/v2/util"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type intersectsQuery struct {
	innerCells []uint64
	crossCells []uint64

	shape index.GeoJSON
	bBox  index.GeoJSON
}

func NewIntersectsQuery(shape index.GeoJSON) Query { _ = "STUB: not implemented"; return *new(Query) }

func (iq *intersectsQuery) Evaluate(geoData segment.GeoShapeV2Data) *util.Bitset {
	_ = "STUB: not implemented"
	return nil
}

func (iq *intersectsQuery) InnerCells() []uint64 { _ = "STUB: not implemented"; return nil }

func (iq *intersectsQuery) CrossCells() []uint64 { _ = "STUB: not implemented"; return nil }
