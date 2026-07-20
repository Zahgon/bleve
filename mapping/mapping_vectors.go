//go:build vectors
// +build vectors

package mapping

import (
	"reflect"
)

var (
	MinVectorDims = 1
	MaxVectorDims = 4096
)

func NewVectorFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func NewVectorBase64FieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func processFlatVector(vecV reflect.Value, dims int) ([]float32, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func processVector(vecI interface{}, dims int) ([]float32, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (fm *FieldMapping) processVector(propertyMightBeVector interface{},
	pathString string, path []string, indexes []uint64, context *walkContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (fm *FieldMapping) processVectorBase64(propertyMightBeVectorBase64 interface{},
	pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func validateFieldMapping(field *FieldMapping, path []string,
	fieldAliasCtx map[string]*FieldMapping) error {
	_ = "STUB: not implemented"
	return nil
}

func validateVectorFieldAlias(field *FieldMapping, path []string,
	fieldAliasCtx map[string]*FieldMapping) error {
	_ = "STUB: not implemented"
	return nil
}

func NormalizeVector(vec []float32) []float32 { _ = "STUB: not implemented"; return nil }

func NormalizeMultiVector(vec []float32, dims int) []float32 { _ = "STUB: not implemented"; return nil }
