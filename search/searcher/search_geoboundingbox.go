package searcher

import (
	"context"

	"github.com/blevesearch/bleve/v2/document"
	"github.com/blevesearch/bleve/v2/geo"
	"github.com/blevesearch/bleve/v2/numeric"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type filterFunc func(key []byte) bool

var (
	GeoBitsShift1       = geo.GeoBits << 1
	GeoBitsShift1Minus1 = GeoBitsShift1 - 1
)

func NewGeoBoundingBoxSearcher(ctx context.Context, indexReader index.IndexReader, minLon, minLat,
	maxLon, maxLat float64, field string, boost float64,
	options search.SearcherOptions, checkBoundaries bool) (
	search.Searcher, error,
) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

var (
	geoMaxShift    = document.GeoPrecisionStep * 4
	geoDetailLevel = ((geo.GeoBits << 1) - geoMaxShift) / 2
)

type closeFunc func() error

func ComputeGeoRange(ctx context.Context, term uint64, shift uint,
	sminLon, sminLat, smaxLon, smaxLat float64, checkBoundaries bool,
	indexReader index.IndexReader, field string) (
	onBoundary [][]byte, notOnBoundary [][]byte, err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func buildIsIndexedFunc(ctx context.Context, indexReader index.IndexReader, field string) (isIndexed filterFunc, closeF closeFunc, err error) {
	_ = "STUB: not implemented"
	return *new(filterFunc), *new(closeFunc), nil
}

func buildRectFilter(ctx context.Context, dvReader index.DocValueReader,
	minLon, minLat, maxLon, maxLat float64,
) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}

type geoRangeCompute struct {
	preallocBytesLen                   int
	preallocBytes                      []byte
	sminLon, sminLat, smaxLon, smaxLat float64
	checkBoundaries                    bool
	onBoundary, notOnBoundary          [][]byte
	isIndexed                          func(term []byte) bool
}

func (grc *geoRangeCompute) makePrefixCoded(in int64, shift uint) (rv numeric.PrefixCoded) {
	_ = "STUB: not implemented"
	return *new(numeric.PrefixCoded)
}

func (grc *geoRangeCompute) computeGeoRange(term uint64, shift uint) {
	_ = "STUB: not implemented"
	return
}

func (grc *geoRangeCompute) relateAndRecurse(start, end uint64, res uint) {
	_ = "STUB: not implemented"
	return
}
