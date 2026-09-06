package geov2

import (
	"github.com/blevesearch/bleve/v2/util"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type Query interface {
	Evaluate(geoData segment.GeoShapeV2Data) *util.Bitset
	InnerCells() []uint64
	CrossCells() []uint64
}

func NewQuery(shape index.GeoJSON, relation string) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
