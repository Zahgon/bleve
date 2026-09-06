package geov2

import segment "github.com/blevesearch/scorch_segment_api/v2"

type queryEvaluator struct {
	queryInnerCells []uint64
	queryCrossCells []uint64

	innerCells  []uint64
	innerDocIDs []uint32

	crossCells  []uint64
	crossDocIDs []uint32
}

func NewQueryEvaluator(query Query, geoData segment.GeoShapeV2Data) *queryEvaluator {
	_ = "STUB: not implemented"
	return nil
}

func binarySearchLeftmostGreaterOrEqual(arr []uint64, target uint64) int {
	_ = "STUB: not implemented"
	return 0
}

func forEachScoredDoc(innerScores, crossScores map[uint32]uint64,
	fn func(id uint32, inner, cross uint64)) {
	_ = "STUB: not implemented"
	return
}

func (qe *queryEvaluator) rangeScanInner(innerScores, crossScores map[uint32]uint64) {
	_ = "STUB: not implemented"
	return
}

func (qe *queryEvaluator) rangeScanCross(innerScores, crossScores map[uint32]uint64) {
	_ = "STUB: not implemented"
	return
}

func rangeScanOne(queryCell uint64, minVal, maxVal, cellLevel uint64,
	indexCells []uint64, docIds []uint32, scores map[uint32]uint64) {
	_ = "STUB: not implemented"
	return
}
