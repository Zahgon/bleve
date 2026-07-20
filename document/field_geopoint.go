package document

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/numeric"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeGeoPointField int

func init() {
	var f GeoPointField
	reflectStaticSizeGeoPointField = int(reflect.TypeOf(f).Size())
}

var GeoPrecisionStep uint = 9

type GeoPointField struct {
	name              string
	arrayPositions    []uint64
	options           index.FieldIndexingOptions
	value             numeric.PrefixCoded
	numPlainTextBytes uint64
	length            int
	frequencies       index.TokenFrequencies

	spatialplugin index.SpatialAnalyzerPlugin
}

func (n *GeoPointField) Size() int { _ = "STUB: not implemented"; return 0 }

func (n *GeoPointField) Name() string { _ = "STUB: not implemented"; return "" }

func (n *GeoPointField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (n *GeoPointField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (n *GeoPointField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (n *GeoPointField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (n *GeoPointField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (n *GeoPointField) Analyze() { _ = "STUB: not implemented"; return }

func (n *GeoPointField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (n *GeoPointField) Lon() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *GeoPointField) Lat() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *GeoPointField) GoString() string { _ = "STUB: not implemented"; return "" }

func (n *GeoPointField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func NewGeoPointFieldFromBytes(name string, arrayPositions []uint64, value []byte) *GeoPointField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoPointField(name string, arrayPositions []uint64, lon, lat float64) *GeoPointField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoPointFieldWithIndexingOptions(name string, arrayPositions []uint64, lon, lat float64, options index.FieldIndexingOptions) *GeoPointField {
	_ = "STUB: not implemented"
	return nil
}

func (n *GeoPointField) SetSpatialAnalyzerPlugin(
	plugin index.SpatialAnalyzerPlugin) {
	_ = "STUB: not implemented"
	return
}
