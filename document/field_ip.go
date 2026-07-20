package document

import (
	"net"
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeIPField int

func init() {
	var f IPField
	reflectStaticSizeIPField = int(reflect.TypeOf(f).Size())
}

const DefaultIPIndexingOptions = index.StoreField | index.IndexField | index.DocValues

type IPField struct {
	name              string
	arrayPositions    []uint64
	options           index.FieldIndexingOptions
	value             net.IP
	numPlainTextBytes uint64
	length            int
	frequencies       index.TokenFrequencies
}

func (b *IPField) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *IPField) Name() string { _ = "STUB: not implemented"; return "" }

func (b *IPField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (b *IPField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (n *IPField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (n *IPField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (n *IPField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (b *IPField) Analyze() { _ = "STUB: not implemented"; return }

func (b *IPField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (b *IPField) IP() (net.IP, error) { _ = "STUB: not implemented"; return *new(net.IP), nil }

func (b *IPField) GoString() string { _ = "STUB: not implemented"; return "" }

func (b *IPField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func NewIPFieldFromBytes(name string, arrayPositions []uint64, value []byte) *IPField {
	_ = "STUB: not implemented"
	return nil
}

func NewIPField(name string, arrayPositions []uint64, v net.IP) *IPField {
	_ = "STUB: not implemented"
	return nil
}

func NewIPFieldWithIndexingOptions(name string, arrayPositions []uint64, b net.IP, options index.FieldIndexingOptions) *IPField {
	_ = "STUB: not implemented"
	return nil
}
