package document

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeBooleanField int

func init() {
	var f BooleanField
	reflectStaticSizeBooleanField = int(reflect.TypeOf(f).Size())
}

const DefaultBooleanIndexingOptions = index.StoreField | index.IndexField | index.DocValues

type BooleanField struct {
	name              string
	arrayPositions    []uint64
	options           index.FieldIndexingOptions
	value             []byte
	numPlainTextBytes uint64
	length            int
	frequencies       index.TokenFrequencies
}

func (b *BooleanField) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *BooleanField) Name() string { _ = "STUB: not implemented"; return "" }

func (b *BooleanField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (b *BooleanField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (b *BooleanField) Analyze() { _ = "STUB: not implemented"; return }

func (b *BooleanField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (b *BooleanField) Boolean() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (b *BooleanField) GoString() string { _ = "STUB: not implemented"; return "" }

func (b *BooleanField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *BooleanField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (b *BooleanField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (b *BooleanField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func NewBooleanFieldFromBytes(name string, arrayPositions []uint64, value []byte) *BooleanField {
	_ = "STUB: not implemented"
	return nil
}

func NewBooleanField(name string, arrayPositions []uint64, b bool) *BooleanField {
	_ = "STUB: not implemented"
	return nil
}

func NewBooleanFieldWithIndexingOptions(name string, arrayPositions []uint64, b bool, options index.FieldIndexingOptions) *BooleanField {
	_ = "STUB: not implemented"
	return nil
}
