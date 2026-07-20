package upsidedown

import (
	"context"
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
	store "github.com/blevesearch/upsidedown_store_api"
)

var reflectStaticSizeIndexReader int

func init() {
	var ir IndexReader
	reflectStaticSizeIndexReader = int(reflect.TypeOf(ir).Size())
}

type IndexReader struct {
	index    *UpsideDownCouch
	kvreader store.KVReader
	docCount uint64
}

func (i *IndexReader) TermFieldReader(ctx context.Context, term []byte, fieldName string, includeFreq, includeNorm, includeTermVectors bool) (index.TermFieldReader, error) {
	_ = "STUB: not implemented"
	return *new(index.TermFieldReader), nil
}

func (i *IndexReader) FieldDict(fieldName string) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *IndexReader) FieldDictRange(fieldName string, startTerm []byte, endTerm []byte) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *IndexReader) FieldDictPrefix(fieldName string, termPrefix []byte) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (i *IndexReader) DocIDReaderAll() (index.DocIDReader, error) {
	_ = "STUB: not implemented"
	return *new(index.DocIDReader), nil
}

func (i *IndexReader) DocIDReaderOnly(ids []string) (index.DocIDReader, error) {
	_ = "STUB: not implemented"
	return *new(index.DocIDReader), nil
}

func (i *IndexReader) Document(id string) (doc index.Document, err error) {
	_ = "STUB: not implemented"
	return *new(index.Document), nil
}

func (i *IndexReader) documentVisitFieldTerms(id index.IndexInternalID, fields []string, visitor index.DocValueVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *IndexReader) Fields() (fields []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexReader) GetInternal(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexReader) DocCount() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *IndexReader) Close() error { _ = "STUB: not implemented"; return nil }

func (i *IndexReader) ExternalID(id index.IndexInternalID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (i *IndexReader) InternalID(id string) (index.IndexInternalID, error) {
	_ = "STUB: not implemented"
	return *new(index.IndexInternalID), nil
}

func incrementBytes(in []byte) []byte { _ = "STUB: not implemented"; return nil }

func (i *IndexReader) DocValueReader(fields []string) (index.DocValueReader, error) {
	_ = "STUB: not implemented"
	return *new(index.DocValueReader), nil
}

type DocValueReader struct {
	i      *IndexReader
	fields []string
}

func (dvr *DocValueReader) VisitDocValues(id index.IndexInternalID,
	visitor index.DocValueVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

func (dvr *DocValueReader) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }
