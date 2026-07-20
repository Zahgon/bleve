package document

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/numeric"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeNumericField int

func init() {
	var f NumericField
	reflectStaticSizeNumericField = int(reflect.TypeOf(f).Size())
}

const DefaultNumericIndexingOptions = index.StoreField | index.IndexField | index.DocValues

const DefaultPrecisionStep uint = 4

type NumericField struct {
	name              string
	arrayPositions    []uint64
	options           index.FieldIndexingOptions
	value             numeric.PrefixCoded
	numPlainTextBytes uint64
	length            int
	frequencies       index.TokenFrequencies
}

func (n *NumericField) Size() int { _ = "STUB: not implemented"; return 0 }

func (n *NumericField) Name() string { _ = "STUB: not implemented"; return "" }

func (n *NumericField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (n *NumericField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (n *NumericField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (n *NumericField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (n *NumericField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (n *NumericField) Analyze() { _ = "STUB: not implemented"; return }

func (n *NumericField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (n *NumericField) Number() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *NumericField) GoString() string { _ = "STUB: not implemented"; return "" }

func (n *NumericField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func NewNumericFieldFromBytes(name string, arrayPositions []uint64, value []byte) *NumericField {
	_ = "STUB: not implemented"
	return nil
}

func NewNumericField(name string, arrayPositions []uint64, number float64) *NumericField {
	_ = "STUB: not implemented"
	return nil
}

func NewNumericFieldWithIndexingOptions(name string, arrayPositions []uint64, number float64, options index.FieldIndexingOptions) *NumericField {
	_ = "STUB: not implemented"
	return nil
}
