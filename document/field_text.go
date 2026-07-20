package document

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/analysis"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeTextField int

func init() {
	var f TextField
	reflectStaticSizeTextField = int(reflect.TypeOf(f).Size())
}

const DefaultTextIndexingOptions = index.IndexField | index.DocValues

type TextField struct {
	name              string
	arrayPositions    []uint64
	options           index.FieldIndexingOptions
	analyzer          analysis.Analyzer
	value             []byte
	numPlainTextBytes uint64
	length            int
	frequencies       index.TokenFrequencies
}

func (t *TextField) Size() int { _ = "STUB: not implemented"; return 0 }

func (t *TextField) Name() string { _ = "STUB: not implemented"; return "" }

func (t *TextField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (t *TextField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (t *TextField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (t *TextField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (t *TextField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (t *TextField) Analyze() { _ = "STUB: not implemented"; return }

func (t *TextField) Analyzer() analysis.Analyzer {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer)
}

func (t *TextField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (t *TextField) Text() string { _ = "STUB: not implemented"; return "" }

func (t *TextField) GoString() string { _ = "STUB: not implemented"; return "" }

func (t *TextField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func NewTextField(name string, arrayPositions []uint64, value []byte) *TextField {
	_ = "STUB: not implemented"
	return nil
}

func NewTextFieldWithIndexingOptions(name string, arrayPositions []uint64, value []byte, options index.FieldIndexingOptions) *TextField {
	_ = "STUB: not implemented"
	return nil
}

func NewTextFieldWithAnalyzer(name string, arrayPositions []uint64, value []byte, analyzer analysis.Analyzer) *TextField {
	_ = "STUB: not implemented"
	return nil
}

func NewTextFieldCustom(name string, arrayPositions []uint64, value []byte, options index.FieldIndexingOptions, analyzer analysis.Analyzer) *TextField {
	_ = "STUB: not implemented"
	return nil
}
