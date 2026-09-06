package scorch

import (
	"github.com/RoaringBitmap/roaring/v2"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"

	zapv11 "github.com/blevesearch/zapx/v11"
	zapv12 "github.com/blevesearch/zapx/v12"
	zapv13 "github.com/blevesearch/zapx/v13"
	zapv14 "github.com/blevesearch/zapx/v14"
	zapv15 "github.com/blevesearch/zapx/v15"
	zapv16 "github.com/blevesearch/zapx/v16"
	zapv17 "github.com/blevesearch/zapx/v17"
)

type SegmentPlugin interface {
	Type() string

	Version() uint32

	New(results []index.Document) (segment.Segment, uint64, error)

	NewUsing(results []index.Document, config map[string]interface{}) (segment.Segment, uint64, error)

	Open(path string) (segment.Segment, error)

	OpenUsing(path string, config map[string]interface{}) (segment.Segment, error)

	Merge(segments []segment.Segment, drops []*roaring.Bitmap, path string,
		closeCh chan struct{}, s segment.StatsReporter) (
		[][]uint64, uint64, error)

	MergeUsing(segments []segment.Segment, drops []*roaring.Bitmap, path string,
		closeCh chan struct{}, s segment.StatsReporter, config map[string]interface{}) (
		[][]uint64, uint64, error)
}

var supportedSegmentPlugins map[string]map[uint32]SegmentPlugin
var defaultSegmentPlugin SegmentPlugin

func init() {
	ResetSegmentPlugins()
	RegisterSegmentPlugin(&zapv17.ZapPlugin{}, true)
	RegisterSegmentPlugin(&zapv16.ZapPlugin{}, false)
	RegisterSegmentPlugin(&zapv15.ZapPlugin{}, false)
	RegisterSegmentPlugin(&zapv14.ZapPlugin{}, false)
	RegisterSegmentPlugin(&zapv13.ZapPlugin{}, false)
	RegisterSegmentPlugin(&zapv12.ZapPlugin{}, false)
	RegisterSegmentPlugin(&zapv11.ZapPlugin{}, false)
}

func ResetSegmentPlugins() { _ = "STUB: not implemented"; return }

func RegisterSegmentPlugin(plugin SegmentPlugin, makeDefault bool) {
	_ = "STUB: not implemented"
	return
}

func SupportedSegmentTypes() (rv []string) { _ = "STUB: not implemented"; return nil }

func SupportedSegmentTypeVersions(typ string) (rv []uint32) { _ = "STUB: not implemented"; return nil }

func chooseSegmentPlugin(forcedSegmentType string,
	forcedSegmentVersion uint32) (SegmentPlugin, error) {
	_ = "STUB: not implemented"
	return *new(SegmentPlugin), nil
}

func (s *Scorch) loadSegmentPlugin(forcedSegmentType string,
	forcedSegmentVersion uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scorch) loadSpatialAnalyzerPlugin(typ string) error { _ = "STUB: not implemented"; return nil }
