package mapping

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/standard"
	"github.com/blevesearch/bleve/v2/analysis/datetime/optional"
	"github.com/blevesearch/bleve/v2/document"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search"
)

var MappingJSONStrict = false

const defaultTypeField = "_type"
const defaultType = "_default"
const defaultField = "_all"
const defaultAnalyzer = standard.Name
const defaultDateTimeParser = optional.Name

type IndexMappingImpl struct {
	TypeMapping           map[string]*DocumentMapping `json:"types,omitempty"`
	DefaultMapping        *DocumentMapping            `json:"default_mapping"`
	TypeField             string                      `json:"type_field"`
	DefaultType           string                      `json:"default_type"`
	DefaultAnalyzer       string                      `json:"default_analyzer"`
	DefaultDateTimeParser string                      `json:"default_datetime_parser"`
	DefaultSynonymSource  string                      `json:"default_synonym_source,omitempty"`
	ScoringModel          string                      `json:"scoring_model,omitempty"`
	DefaultField          string                      `json:"default_field"`
	StoreDynamic          bool                        `json:"store_dynamic"`
	IndexDynamic          bool                        `json:"index_dynamic"`
	DocValuesDynamic      bool                        `json:"docvalues_dynamic"`
	CustomAnalysis        *customAnalysis             `json:"analysis,omitempty"`
	cache                 *registry.Cache
}

func (im *IndexMappingImpl) AddCustomCharFilter(name string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) AddCustomTokenizer(name string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) AddCustomTokenMap(name string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) AddCustomTokenFilter(name string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) AddCustomAnalyzer(name string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) AddCustomDateTimeParser(name string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) AddSynonymSource(name string, config map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexMapping() *IndexMappingImpl { _ = "STUB: not implemented"; return nil }

func (im *IndexMappingImpl) Validate() error { _ = "STUB: not implemented"; return nil }

func (im *IndexMappingImpl) AddDocumentMapping(doctype string, dm *DocumentMapping) {
	_ = "STUB: not implemented"
	return
}

func (im *IndexMappingImpl) mappingForType(docType string) *DocumentMapping {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (im *IndexMappingImpl) determineType(data interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (im *IndexMappingImpl) MapDocument(doc *document.Document, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) MapSynonymDocument(doc *document.Document, collection string, input []string, synonyms []string) error {
	_ = "STUB: not implemented"
	return nil
}

type walkContext struct {
	doc             *document.Document
	im              *IndexMappingImpl
	dm              *DocumentMapping
	excludedFromAll []string
}

func (im *IndexMappingImpl) newWalkContext(doc *document.Document, dm *DocumentMapping) *walkContext {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) AnalyzerNameForPath(path string) string {
	_ = "STUB: not implemented"
	return ""
}

func (im *IndexMappingImpl) AnalyzerNamed(name string) analysis.Analyzer {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer)
}

func (im *IndexMappingImpl) DateTimeParserNamed(name string) analysis.DateTimeParser {
	_ = "STUB: not implemented"
	return *new(analysis.DateTimeParser)
}

func (im *IndexMappingImpl) AnalyzeText(analyzerName string, text []byte) (analysis.TokenStream, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream), nil
}

func (im *IndexMappingImpl) FieldAnalyzer(field string) string {
	_ = "STUB: not implemented"
	return ""
}

func (im *IndexMappingImpl) FieldMappingForPath(path string) FieldMapping {
	_ = "STUB: not implemented"
	return *new(FieldMapping)
}

func (im *IndexMappingImpl) DefaultSearchField() string { _ = "STUB: not implemented"; return "" }

func (im *IndexMappingImpl) SynonymSourceNamed(name string) analysis.SynonymSource {
	_ = "STUB: not implemented"
	return *new(analysis.SynonymSource)
}

func (im *IndexMappingImpl) SynonymSourceForPath(path string) string {
	_ = "STUB: not implemented"
	return ""
}

func (im *IndexMappingImpl) SynonymCount() int { _ = "STUB: not implemented"; return 0 }

func (im *IndexMappingImpl) SynonymSourceVisitor(visitor analysis.SynonymSourceVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) buildNestedPrefixes() map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func (im *IndexMappingImpl) NestedDepth(fs search.FieldSet) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (im *IndexMappingImpl) CountNested() int { _ = "STUB: not implemented"; return 0 }

func (im *IndexMappingImpl) IntersectsPrefix(fs search.FieldSet) bool {
	_ = "STUB: not implemented"
	return false
}
