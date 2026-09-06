package document

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
	"github.com/blevesearch/geo/geojson"
)

var reflectStaticSizeGeoShapeV2Field int

func init() {
	var f GeoShapeV2Field
	reflectStaticSizeGeoShapeV2Field = int(reflect.TypeOf(f).Size())
}

type GeoShapeV2Field struct {
	name string

	shape      index.GeoJSON
	inner      []uint64
	cross      []uint64
	scoreInner uint64
	scoreCross uint64
	bBoxBytes  []byte
	shapeBytes []byte

	options index.FieldIndexingOptions
}

func (f *GeoShapeV2Field) Name() string { _ = "STUB: not implemented"; return "" }

func (f *GeoShapeV2Field) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (f *GeoShapeV2Field) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (f *GeoShapeV2Field) Analyze() { _ = "STUB: not implemented"; return }

func (f *GeoShapeV2Field) Value() []byte { _ = "STUB: not implemented"; return nil }

func (f *GeoShapeV2Field) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *GeoShapeV2Field) Size() int { _ = "STUB: not implemented"; return 0 }

func (f *GeoShapeV2Field) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (f *GeoShapeV2Field) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (f *GeoShapeV2Field) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (f *GeoShapeV2Field) InnerCells() []uint64 { _ = "STUB: not implemented"; return nil }

func (f *GeoShapeV2Field) CrossCells() []uint64 { _ = "STUB: not implemented"; return nil }

func (f *GeoShapeV2Field) EncodedBoundingBox() []byte { _ = "STUB: not implemented"; return nil }

func (f *GeoShapeV2Field) EncodedShape() []byte { _ = "STUB: not implemented"; return nil }

func (f *GeoShapeV2Field) Scores() (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

func NewGeoShapeV2FieldFromShapeWithIndexingOptions(name string, geoShape *geojson.GeoShape,
	options index.FieldIndexingOptions) *GeoShapeV2Field {
	_ = "STUB: not implemented"
	return nil
}

func NewGeometryCollectionV2FieldFromShapesWithIndexingOptions(name string,
	geoShapes []*geojson.GeoShape, options index.FieldIndexingOptions) *GeoShapeV2Field {
	_ = "STUB: not implemented"
	return nil
}
