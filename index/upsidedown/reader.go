package upsidedown

import (
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
	store "github.com/blevesearch/upsidedown_store_api"
)

var reflectStaticSizeUpsideDownCouchTermFieldReader int
var reflectStaticSizeUpsideDownCouchDocIDReader int

func init() {
	var tfr UpsideDownCouchTermFieldReader
	reflectStaticSizeUpsideDownCouchTermFieldReader =
		int(reflect.TypeOf(tfr).Size())
	var cdr UpsideDownCouchDocIDReader
	reflectStaticSizeUpsideDownCouchDocIDReader =
		int(reflect.TypeOf(cdr).Size())
}

type UpsideDownCouchTermFieldReader struct {
	count              uint64
	indexReader        *IndexReader
	iterator           store.KVIterator
	term               []byte
	tfrNext            *TermFrequencyRow
	tfrPrealloc        TermFrequencyRow
	keyBuf             []byte
	field              uint16
	includeTermVectors bool
}

func (r *UpsideDownCouchTermFieldReader) Size() int { _ = "STUB: not implemented"; return 0 }

func newUpsideDownCouchTermFieldReader(indexReader *IndexReader, term []byte, field uint16, includeFreq, includeNorm, includeTermVectors bool) (*UpsideDownCouchTermFieldReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpsideDownCouchTermFieldReader) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *UpsideDownCouchTermFieldReader) Next(preAlloced *index.TermFieldDoc) (*index.TermFieldDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpsideDownCouchTermFieldReader) Advance(docID index.IndexInternalID, preAlloced *index.TermFieldDoc) (rv *index.TermFieldDoc, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpsideDownCouchTermFieldReader) Close() error { _ = "STUB: not implemented"; return nil }

type UpsideDownCouchDocIDReader struct {
	indexReader *IndexReader
	iterator    store.KVIterator
	only        []string
	onlyPos     int
	onlyMode    bool
}

func (r *UpsideDownCouchDocIDReader) Size() int { _ = "STUB: not implemented"; return 0 }

func newUpsideDownCouchDocIDReader(indexReader *IndexReader) (*UpsideDownCouchDocIDReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newUpsideDownCouchDocIDReaderOnly(indexReader *IndexReader, ids []string) (*UpsideDownCouchDocIDReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpsideDownCouchDocIDReader) Next() (index.IndexInternalID, error) {
	_ = "STUB: not implemented"
	return *new(index.IndexInternalID), nil
}

func (r *UpsideDownCouchDocIDReader) Advance(docID index.IndexInternalID) (index.IndexInternalID, error) {
	_ = "STUB: not implemented"
	return *new(index.IndexInternalID), nil
}

func (r *UpsideDownCouchDocIDReader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *UpsideDownCouchDocIDReader) nextOnly() bool { _ = "STUB: not implemented"; return false }
