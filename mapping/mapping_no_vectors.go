//go:build !vectors
// +build !vectors

package mapping

func NewVectorFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func NewVectorBase64FieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }

func (fm *FieldMapping) processVector(propertyMightBeVector interface{},
	pathString string, path []string, indexes []uint64, context *walkContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (fm *FieldMapping) processVectorBase64(propertyMightBeVector interface{},
	pathString string, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func validateFieldMapping(field *FieldMapping, path []string,
	fieldAliasCtx map[string]*FieldMapping) error {
	_ = "STUB: not implemented"
	return nil
}
