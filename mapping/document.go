package mapping

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/registry"
)

type DocumentMapping struct {
	Enabled              bool                        `json:"enabled"`
	Dynamic              bool                        `json:"dynamic"`
	Properties           map[string]*DocumentMapping `json:"properties,omitempty"`
	Fields               []*FieldMapping             `json:"fields,omitempty"`
	Nested               bool                        `json:"nested,omitempty"`
	DefaultAnalyzer      string                      `json:"default_analyzer,omitempty"`
	DefaultSynonymSource string                      `json:"default_synonym_source,omitempty"`

	StructTagKey string `json:"struct_tag_key,omitempty"`
}

func (dm *DocumentMapping) Validate(cache *registry.Cache,
	path []string, fieldAliasCtx map[string]*FieldMapping,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFieldType(field *FieldMapping) error { _ = "STUB: not implemented"; return nil }

func (dm *DocumentMapping) analyzerNameForPath(path string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *DocumentMapping) synonymSourceForPath(path string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *DocumentMapping) fieldDescribedByPath(path string) *FieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (dm *DocumentMapping) documentMappingForPathElements(pathElements []string) (
	*DocumentMapping, *DocumentMapping,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DocumentMapping) documentMappingForPath(path string) (
	*DocumentMapping, *DocumentMapping,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDocumentMapping() *DocumentMapping { _ = "STUB: not implemented"; return nil }

func NewNestedDocumentMapping() *DocumentMapping { _ = "STUB: not implemented"; return nil }

func NewDocumentStaticMapping() *DocumentMapping { _ = "STUB: not implemented"; return nil }

func NewNestedDocumentStaticMapping() *DocumentMapping { _ = "STUB: not implemented"; return nil }

func NewDocumentDisabledMapping() *DocumentMapping { _ = "STUB: not implemented"; return nil }

func (dm *DocumentMapping) AddSubDocumentMapping(property string, sdm *DocumentMapping) {
	_ = "STUB: not implemented"
	return
}

func (dm *DocumentMapping) AddFieldMappingsAt(property string, fms ...*FieldMapping) {
	_ = "STUB: not implemented"
	return
}

func (dm *DocumentMapping) AddFieldMapping(fm *FieldMapping) { _ = "STUB: not implemented"; return }

func (dm *DocumentMapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (dm *DocumentMapping) defaultAnalyzerName(path []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (dm *DocumentMapping) defaultSynonymSource(path []string) string {
	_ = "STUB: not implemented"
	return ""
}

func baseType(v interface{}) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (dm *DocumentMapping) walkDocument(data interface{}, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}

func (dm *DocumentMapping) processProperty(property interface{}, path []string, indexes []uint64, context *walkContext) {
	_ = "STUB: not implemented"
	return
}
