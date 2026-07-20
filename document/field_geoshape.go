package document

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
	"github.com/blevesearch/geo/geojson"
)

var reflectStaticSizeGeoShapeField int

func init() {
	var f GeoShapeField
	reflectStaticSizeGeoShapeField = int(reflect.TypeOf(f).Size())
}

const DefaultGeoShapeIndexingOptions = index.IndexField | index.DocValues

type GeoShapeField struct {
	name              string
	shape             index.GeoJSON
	arrayPositions    []uint64
	options           index.FieldIndexingOptions
	numPlainTextBytes uint64
	length            int
	encodedValue      []byte
	value             []byte

	frequencies index.TokenFrequencies
}

func (n *GeoShapeField) Size() int { _ = "STUB: not implemented"; return 0 }

func (n *GeoShapeField) Name() string { _ = "STUB: not implemented"; return "" }

func (n *GeoShapeField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (n *GeoShapeField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (n *GeoShapeField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (n *GeoShapeField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (n *GeoShapeField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (n *GeoShapeField) Analyze() { _ = "STUB: not implemented"; return }

func (n *GeoShapeField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (n *GeoShapeField) GoString() string { _ = "STUB: not implemented"; return "" }

func (n *GeoShapeField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *GeoShapeField) EncodedShape() []byte { _ = "STUB: not implemented"; return nil }

func NewGeoShapeField(name string, arrayPositions []uint64,
	coordinates [][][][]float64, typ string) *GeoShapeField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoShapeFieldFromBytes(name string, arrayPositions []uint64,
	value []byte) *GeoShapeField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoShapeFieldWithIndexingOptions(name string, arrayPositions []uint64,
	coordinates [][][][]float64, typ string,
	options index.FieldIndexingOptions) *GeoShapeField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoShapeFieldFromShapeWithIndexingOptions(name string, arrayPositions []uint64,
	geoShape *geojson.GeoShape, options index.FieldIndexingOptions) *GeoShapeField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeometryCollectionFieldWithIndexingOptions(name string,
	arrayPositions []uint64, coordinates [][][][][]float64, types []string,
	options index.FieldIndexingOptions) *GeoShapeField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeometryCollectionFieldFromShapesWithIndexingOptions(name string,
	arrayPositions []uint64, geoShapes []*geojson.GeoShape,
	options index.FieldIndexingOptions) *GeoShapeField {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoCircleFieldWithIndexingOptions(name string, arrayPositions []uint64,
	centerPoint []float64, radius string,
	options index.FieldIndexingOptions) *GeoShapeField {
	_ = "STUB: not implemented"
	return nil
}

func (n *GeoShapeField) GeoShape() (index.GeoJSON, error) {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON), nil
}
