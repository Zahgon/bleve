//go:build vectors
// +build vectors

package document

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeVectorBase64Field int

func init() {
	var f VectorBase64Field
	reflectStaticSizeVectorBase64Field = int(reflect.TypeOf(f).Size())
}

type VectorBase64Field struct {
	vectorField    *VectorField
	base64Encoding string
}

func (n *VectorBase64Field) Size() int { _ = "STUB: not implemented"; return 0 }

func (n *VectorBase64Field) Name() string { _ = "STUB: not implemented"; return "" }

func (n *VectorBase64Field) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (n *VectorBase64Field) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (n *VectorBase64Field) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *VectorBase64Field) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (n *VectorBase64Field) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (n *VectorBase64Field) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (n *VectorBase64Field) Analyze() { _ = "STUB: not implemented"; return }

func (n *VectorBase64Field) Value() []byte { _ = "STUB: not implemented"; return nil }

func (n *VectorBase64Field) GoString() string { _ = "STUB: not implemented"; return "" }

func NewVectorBase64Field(name string, arrayPositions []uint64, vectorBase64 string,
	dims int, similarity, vectorIndexOptimizedFor string) (*VectorBase64Field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecodeVector(encodedValue string) ([]float32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *VectorBase64Field) Vector() []float32 { _ = "STUB: not implemented"; return nil }

func (n *VectorBase64Field) Dims() int { _ = "STUB: not implemented"; return 0 }

func (n *VectorBase64Field) Similarity() string { _ = "STUB: not implemented"; return "" }

func (n *VectorBase64Field) IndexOptimizedFor() string { _ = "STUB: not implemented"; return "" }
