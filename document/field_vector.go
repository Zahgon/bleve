//go:build vectors
// +build vectors

package document

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeVectorField int

func init() {
	var f VectorField
	reflectStaticSizeVectorField = int(reflect.TypeOf(f).Size())
}

const DefaultVectorIndexingOptions = index.IndexField

type VectorField struct {
	name                    string
	dims                    int
	similarity              string
	options                 index.FieldIndexingOptions
	value                   []float32
	numPlainTextBytes       uint64
	vectorIndexOptimizedFor string
}

func (n *VectorField) Size() int { _ = "STUB: not implemented"; return 0 }

func (n *VectorField) Name() string { _ = "STUB: not implemented"; return "" }

func (n *VectorField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (n *VectorField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (n *VectorField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *VectorField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (n *VectorField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (n *VectorField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (n *VectorField) Analyze() { _ = "STUB: not implemented"; return }

func (n *VectorField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (n *VectorField) GoString() string { _ = "STUB: not implemented"; return "" }

func NewVectorField(name string, arrayPositions []uint64,
	vector []float32, dims int, similarity, vectorIndexOptimizedFor string) *VectorField {
	_ = "STUB: not implemented"
	return nil
}

func NewVectorFieldWithIndexingOptions(name string, arrayPositions []uint64,
	vector []float32, dims int, similarity, vectorIndexOptimizedFor string,
	options index.FieldIndexingOptions) *VectorField {
	_ = "STUB: not implemented"
	return nil
}

func numBytesFloat32s(value []float32) uint64 { _ = "STUB: not implemented"; return 0 }

func (n *VectorField) Vector() []float32 { _ = "STUB: not implemented"; return nil }

func (n *VectorField) Dims() int { _ = "STUB: not implemented"; return 0 }

func (n *VectorField) Similarity() string { _ = "STUB: not implemented"; return "" }

func (n *VectorField) IndexOptimizedFor() string { _ = "STUB: not implemented"; return "" }
