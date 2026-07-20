package document

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeDocument int

func init() {
	var d Document
	reflectStaticSizeDocument = int(reflect.TypeOf(d).Size())
}

type Document struct {
	id               string
	Fields           []Field     `json:"fields"`
	NestedDocuments  []*Document `json:"nested_documents"`
	CompositeFields  []*CompositeField
	StoredFieldsSize uint64
	indexed          bool
}

func (d *Document) StoredFieldsBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func NewDocument(id string) *Document { _ = "STUB: not implemented"; return nil }

func NewSynonymDocument(id string) *Document { _ = "STUB: not implemented"; return nil }

func (d *Document) Size() int { _ = "STUB: not implemented"; return 0 }

func (d *Document) AddField(f Field) *Document { _ = "STUB: not implemented"; return nil }

func (d *Document) GoString() string { _ = "STUB: not implemented"; return "" }

func (d *Document) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *Document) ID() string { _ = "STUB: not implemented"; return "" }

func (d *Document) SetID(id string) { _ = "STUB: not implemented"; return }

func (d *Document) AddIDField() { _ = "STUB: not implemented"; return }

func (d *Document) VisitFields(visitor index.FieldVisitor) { _ = "STUB: not implemented"; return }

func (d *Document) VisitComposite(visitor index.CompositeFieldVisitor) {
	_ = "STUB: not implemented"
	return
}

func (d *Document) HasComposite() bool { _ = "STUB: not implemented"; return false }

func (d *Document) VisitSynonymFields(visitor index.SynonymFieldVisitor) {
	_ = "STUB: not implemented"
	return
}

func (d *Document) SetIndexed() { _ = "STUB: not implemented"; return }

func (d *Document) Indexed() bool { _ = "STUB: not implemented"; return false }

func (d *Document) AddNestedDocument(doc *Document) { _ = "STUB: not implemented"; return }

func (d *Document) VisitNestedDocuments(visitor func(doc index.Document)) {
	_ = "STUB: not implemented"
	return
}
