package document

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeCompositeField int

func init() {
	var cf CompositeField
	reflectStaticSizeCompositeField = int(reflect.TypeOf(cf).Size())
}

const DefaultCompositeIndexingOptions = index.IndexField

type CompositeField struct {
	name                 string
	includedFields       map[string]bool
	excludedFields       map[string]bool
	defaultInclude       bool
	options              index.FieldIndexingOptions
	totalLength          int
	compositeFrequencies index.TokenFrequencies
}

func NewCompositeField(name string, defaultInclude bool, include []string, exclude []string) *CompositeField {
	_ = "STUB: not implemented"
	return nil
}

func NewCompositeFieldWithIndexingOptions(name string, defaultInclude bool, include []string, exclude []string, options index.FieldIndexingOptions) *CompositeField {
	_ = "STUB: not implemented"
	return nil
}

func (c *CompositeField) Size() int { _ = "STUB: not implemented"; return 0 }

func (c *CompositeField) Name() string { _ = "STUB: not implemented"; return "" }

func (c *CompositeField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (c *CompositeField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (c *CompositeField) Analyze() { _ = "STUB: not implemented"; return }

func (c *CompositeField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (c *CompositeField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *CompositeField) includesField(field string) bool { _ = "STUB: not implemented"; return false }

func (c *CompositeField) Compose(field string, length int, freq index.TokenFrequencies) {
	_ = "STUB: not implemented"
	return
}

func (c *CompositeField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (c *CompositeField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (c *CompositeField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}
