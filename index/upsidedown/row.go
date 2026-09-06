package upsidedown

import (
	"encoding/binary"
	"reflect"
)

var (
	reflectStaticSizeTermFrequencyRow int
	reflectStaticSizeTermVector       int
)

func init() {
	var tfr TermFrequencyRow
	reflectStaticSizeTermFrequencyRow = int(reflect.TypeOf(tfr).Size())
	var tv TermVector
	reflectStaticSizeTermVector = int(reflect.TypeOf(tv).Size())
}

const ByteSeparator byte = 0xff

type UpsideDownCouchRowStream chan UpsideDownCouchRow

type UpsideDownCouchRow interface {
	KeySize() int
	KeyTo([]byte) (int, error)
	Key() []byte
	Value() []byte
	ValueSize() int
	ValueTo([]byte) (int, error)
}

func ParseFromKeyValue(key, value []byte) (UpsideDownCouchRow, error) {
	_ = "STUB: not implemented"
	return *new(UpsideDownCouchRow), nil
}

type VersionRow struct {
	version uint8
}

func (v *VersionRow) Key() []byte { _ = "STUB: not implemented"; return nil }

func (v *VersionRow) KeySize() int { _ = "STUB: not implemented"; return 0 }

func (v *VersionRow) KeyTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *VersionRow) Value() []byte { _ = "STUB: not implemented"; return nil }

func (v *VersionRow) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (v *VersionRow) ValueTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *VersionRow) String() string { _ = "STUB: not implemented"; return "" }

func NewVersionRow(version uint8) *VersionRow { _ = "STUB: not implemented"; return nil }

func NewVersionRowKV(key, value []byte) (*VersionRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type InternalRow struct {
	key []byte
	val []byte
}

func (i *InternalRow) Key() []byte { _ = "STUB: not implemented"; return nil }

func (i *InternalRow) KeySize() int { _ = "STUB: not implemented"; return 0 }

func (i *InternalRow) KeyTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *InternalRow) Value() []byte { _ = "STUB: not implemented"; return nil }

func (i *InternalRow) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (i *InternalRow) ValueTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *InternalRow) String() string { _ = "STUB: not implemented"; return "" }

func NewInternalRow(key, val []byte) *InternalRow { _ = "STUB: not implemented"; return nil }

func NewInternalRowKV(key, value []byte) (*InternalRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FieldRow struct {
	index uint16
	name  string
}

func (f *FieldRow) Key() []byte { _ = "STUB: not implemented"; return nil }

func (f *FieldRow) KeySize() int { _ = "STUB: not implemented"; return 0 }

func (f *FieldRow) KeyTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *FieldRow) Value() []byte { _ = "STUB: not implemented"; return nil }

func (f *FieldRow) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (f *FieldRow) ValueTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *FieldRow) String() string { _ = "STUB: not implemented"; return "" }

func NewFieldRow(index uint16, name string) *FieldRow { _ = "STUB: not implemented"; return nil }

func NewFieldRowKV(key, value []byte) (*FieldRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const DictionaryRowMaxValueSize = binary.MaxVarintLen64

type DictionaryRow struct {
	term  []byte
	count uint64
	field uint16
}

func (dr *DictionaryRow) Key() []byte { _ = "STUB: not implemented"; return nil }

func (dr *DictionaryRow) KeySize() int { _ = "STUB: not implemented"; return 0 }

func dictionaryRowKeySize(term []byte) int { _ = "STUB: not implemented"; return 0 }

func (dr *DictionaryRow) KeyTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func dictionaryRowKeyTo(buf []byte, field uint16, term []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func (dr *DictionaryRow) Value() []byte { _ = "STUB: not implemented"; return nil }

func (dr *DictionaryRow) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (dr *DictionaryRow) ValueTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (dr *DictionaryRow) String() string { _ = "STUB: not implemented"; return "" }

func NewDictionaryRow(term []byte, field uint16, count uint64) *DictionaryRow {
	_ = "STUB: not implemented"
	return nil
}

func NewDictionaryRowKV(key, value []byte) (*DictionaryRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDictionaryRowK(key []byte) (*DictionaryRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr *DictionaryRow) parseDictionaryK(key []byte) error { _ = "STUB: not implemented"; return nil }

func (dr *DictionaryRow) parseDictionaryV(value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func dictionaryRowParseV(value []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

type TermVector struct {
	field          uint16
	arrayPositions []uint64
	pos            uint64
	start          uint64
	end            uint64
}

func (tv *TermVector) Size() int { _ = "STUB: not implemented"; return 0 }

func (tv *TermVector) String() string { _ = "STUB: not implemented"; return "" }

type TermFrequencyRow struct {
	term    []byte
	doc     []byte
	freq    uint64
	vectors []*TermVector
	norm    float32
	field   uint16
}

func (tfr *TermFrequencyRow) Size() int { _ = "STUB: not implemented"; return 0 }

func (tfr *TermFrequencyRow) Term() []byte { _ = "STUB: not implemented"; return nil }

func (tfr *TermFrequencyRow) Freq() uint64 { _ = "STUB: not implemented"; return 0 }

func (tfr *TermFrequencyRow) ScanPrefixForField() []byte { _ = "STUB: not implemented"; return nil }

func (tfr *TermFrequencyRow) ScanPrefixForFieldTermPrefix() []byte {
	_ = "STUB: not implemented"
	return nil
}

func (tfr *TermFrequencyRow) ScanPrefixForFieldTerm() []byte { _ = "STUB: not implemented"; return nil }

func (tfr *TermFrequencyRow) Key() []byte { _ = "STUB: not implemented"; return nil }

func (tfr *TermFrequencyRow) KeySize() int { _ = "STUB: not implemented"; return 0 }

func termFrequencyRowKeySize(term, doc []byte) int { _ = "STUB: not implemented"; return 0 }

func (tfr *TermFrequencyRow) KeyTo(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func termFrequencyRowKeyTo(buf []byte, field uint16, term, doc []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func (tfr *TermFrequencyRow) KeyAppendTo(buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tfr *TermFrequencyRow) DictionaryRowKey() []byte { _ = "STUB: not implemented"; return nil }

func (tfr *TermFrequencyRow) DictionaryRowKeySize() int { _ = "STUB: not implemented"; return 0 }

func (tfr *TermFrequencyRow) DictionaryRowKeyTo(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tfr *TermFrequencyRow) Value() []byte { _ = "STUB: not implemented"; return nil }

func (tfr *TermFrequencyRow) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (tfr *TermFrequencyRow) ValueTo(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tfr *TermFrequencyRow) String() string { _ = "STUB: not implemented"; return "" }

func InitTermFrequencyRow(tfr *TermFrequencyRow, term []byte, field uint16, docID []byte, freq uint64, norm float32) *TermFrequencyRow {
	_ = "STUB: not implemented"
	return nil
}

func NewTermFrequencyRow(term []byte, field uint16, docID []byte, freq uint64, norm float32) *TermFrequencyRow {
	_ = "STUB: not implemented"
	return nil
}

func NewTermFrequencyRowWithTermVectors(term []byte, field uint16, docID []byte, freq uint64, norm float32, vectors []*TermVector) *TermFrequencyRow {
	_ = "STUB: not implemented"
	return nil
}

func NewTermFrequencyRowK(key []byte) (*TermFrequencyRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tfr *TermFrequencyRow) parseK(key []byte) error { _ = "STUB: not implemented"; return nil }

func (tfr *TermFrequencyRow) parseKDoc(key []byte, term []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (tfr *TermFrequencyRow) parseV(value []byte, includeTermVectors bool) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTermFrequencyRowKV(key, value []byte) (*TermFrequencyRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BackIndexRow struct {
	doc           []byte
	termsEntries  []*BackIndexTermsEntry
	storedEntries []*BackIndexStoreEntry
}

func (br *BackIndexRow) AllTermKeys() [][]byte { _ = "STUB: not implemented"; return nil }

func (br *BackIndexRow) AllStoredKeys() [][]byte { _ = "STUB: not implemented"; return nil }

func (br *BackIndexRow) Key() []byte { _ = "STUB: not implemented"; return nil }

func (br *BackIndexRow) KeySize() int { _ = "STUB: not implemented"; return 0 }

func (br *BackIndexRow) KeyTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (br *BackIndexRow) Value() []byte { _ = "STUB: not implemented"; return nil }

func (br *BackIndexRow) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (br *BackIndexRow) ValueTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (br *BackIndexRow) String() string { _ = "STUB: not implemented"; return "" }

func NewBackIndexRow(docID []byte, entries []*BackIndexTermsEntry, storedFields []*BackIndexStoreEntry) *BackIndexRow {
	_ = "STUB: not implemented"
	return nil
}

func NewBackIndexRowKV(key, value []byte) (*BackIndexRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type StoredRow struct {
	doc            []byte
	field          uint16
	arrayPositions []uint64
	typ            byte
	value          []byte
}

func (s *StoredRow) Key() []byte { _ = "STUB: not implemented"; return nil }

func (s *StoredRow) KeySize() int { _ = "STUB: not implemented"; return 0 }

func (s *StoredRow) KeyTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *StoredRow) Value() []byte { _ = "STUB: not implemented"; return nil }

func (s *StoredRow) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (s *StoredRow) ValueTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *StoredRow) String() string { _ = "STUB: not implemented"; return "" }

func (s *StoredRow) ScanPrefixForDoc() []byte { _ = "STUB: not implemented"; return nil }

func NewStoredRow(docID []byte, field uint16, arrayPositions []uint64, typ byte, value []byte) *StoredRow {
	_ = "STUB: not implemented"
	return nil
}

func NewStoredRowK(key []byte) (*StoredRow, error) { _ = "STUB: not implemented"; return nil, nil }

func NewStoredRowKV(key, value []byte) (*StoredRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type backIndexFieldTermVisitor func(field uint32, term []byte)

func visitBackIndexRow(data []byte, callback backIndexFieldTermVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitBackIndexRowFieldTerms(data []byte, callback backIndexFieldTermVisitor) error {
	_ = "STUB: not implemented"
	return nil
}
