package mapping

import (
	"net"
	"time"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/document"
	index "github.com/blevesearch/bleve_index_api"
	"github.com/blevesearch/geo/geojson"
)

var (
	IndexDynamic     = true
	StoreDynamic     = true
	DocValuesDynamic = true
)

type FieldMapping struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	Analyzer string `json:"analyzer,omitempty"`

	Store bool `json:"store,omitempty"`
	Index bool `json:"index,omitempty"`

	IncludeTermVectors bool   `json:"include_term_vectors,omitempty"`
	IncludeInAll       bool   `json:"include_in_all,omitempty"`
	DateFormat         string `json:"date_format,omitempty"`

	DocValues bool `json:"docvalues,omitempty"`

	SkipFreqNorm bool `json:"skip_freq_norm,omitempty"`

	Dims int `json:"dims,omitempty"`

	Similarity string `json:"similarity,omitempty"`

	VectorIndexOptimizedFor string `json:"vector_index_optimized_for,omitempty"`

	SynonymSource string `json:"synonym_source,omitempty"`

	GPU bool `json:"gpu,omitempty"`
}

func NewTextFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func newTextFieldMappingDynamic(im *IndexMappingImpl) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func NewKeywordFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func NewNumericFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func newNumericFieldMappingDynamic(im *IndexMappingImpl) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func NewDateTimeFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func newDateTimeFieldMappingDynamic(im *IndexMappingImpl) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func NewBooleanFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func newBooleanFieldMappingDynamic(im *IndexMappingImpl) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoPointFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func NewGeoShapeFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func NewGeoShapeV2FieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func NewIPFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func (fm *FieldMapping) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (fm *FieldMapping) processString(propertyValueString string, pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processFloat64(propertyValFloat float64, pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processTime(propertyValueTime time.Time, layout string, pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processBoolean(propertyValueBool bool, pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processGeoPoint(propertyMightBeGeoPoint interface{}, pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processIP(ip net.IP, pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processGeoShape(propertyMightBeGeoShape interface{},
	pathString string, path []string, indexes []uint64, context *walkContext,
) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processGeoShapeV2(propertyMightBeGeoShape interface{},
	pathString string, path []string, context *walkContext,
) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) processGeoShapeInternal(
	propertyMightBeGeoShape interface{},
	pathString string, path []string, context *walkContext,
	makeCollection func(fieldName string, shapes []*geojson.GeoShape,
		options index.FieldIndexingOptions) document.Field,
	makeShape func(fieldName string, shape *geojson.GeoShape,
		options index.FieldIndexingOptions) document.Field,
) {
	_ = "STUB: not implemented"
	return
}

func (fm *FieldMapping) analyzerForField(path []string, context *walkContext) analysis.Analyzer {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer)
}

func getFieldName(pathString string, path []string, fieldMapping *FieldMapping) string {
	_ = "STUB: not implemented"
	return ""
}

func (fm *FieldMapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
