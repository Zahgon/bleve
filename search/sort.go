package search

import (
	"encoding/json"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/blevesearch/bleve/v2/numeric"
)

var (
	HighTerm = strings.Repeat(string(utf8.MaxRune), 3)
	LowTerm  = string([]byte{0x00})
)

type SearchSort interface {
	UpdateVisitor(field string, term []byte)
	Value(a *DocumentMatch) string
	DecodeValue(value string) string
	Descending() bool

	RequiresDocID() bool
	RequiresScoring() bool
	RequiresFields() []string

	Reverse()

	Copy() SearchSort
}

func ParseSearchSortObj(input map[string]interface{}) (SearchSort, error) {
	_ = "STUB: not implemented"
	return *new(SearchSort), nil
}

func ParseSearchSortString(input string) SearchSort {
	_ = "STUB: not implemented"
	return *new(SearchSort)
}

func ParseSearchSortJSON(input json.RawMessage) (SearchSort, error) {
	_ = "STUB: not implemented"
	return *new(SearchSort), nil
}

func ParseSortOrderStrings(in []string) SortOrder {
	_ = "STUB: not implemented"
	return *new(SortOrder)
}

func ParseSortOrderJSON(in []json.RawMessage) (SortOrder, error) {
	_ = "STUB: not implemented"
	return *new(SortOrder), nil
}

type SortOrder []SearchSort

func (so SortOrder) Value(doc *DocumentMatch) { _ = "STUB: not implemented"; return }

func (so SortOrder) UpdateVisitor(field string, term []byte) { _ = "STUB: not implemented"; return }

func (so SortOrder) Copy() SortOrder { _ = "STUB: not implemented"; return *new(SortOrder) }

func (so SortOrder) Compare(cachedScoring, cachedDesc []bool, i, j *DocumentMatch) int {
	_ = "STUB: not implemented"
	return 0
}

func (so SortOrder) RequiresScore() bool { _ = "STUB: not implemented"; return false }

func (so SortOrder) RequiresDocID() bool { _ = "STUB: not implemented"; return false }

func (so SortOrder) RequiredFields() []string { _ = "STUB: not implemented"; return nil }

func (so SortOrder) CacheIsScore() []bool { _ = "STUB: not implemented"; return nil }

func (so SortOrder) CacheDescending() []bool { _ = "STUB: not implemented"; return nil }

func (so SortOrder) Reverse() { _ = "STUB: not implemented"; return }

type SortFieldType int

const (
	SortFieldAuto SortFieldType = iota

	SortFieldAsString

	SortFieldAsNumber

	SortFieldAsDate
)

type SortFieldMode int

const (
	SortFieldDefault SortFieldMode = iota

	SortFieldMin

	SortFieldMax
)

type SortFieldMissing int

const (
	SortFieldMissingLast SortFieldMissing = iota

	SortFieldMissingFirst
)

type SortField struct {
	Field   string
	Desc    bool
	Type    SortFieldType
	Mode    SortFieldMode
	Missing SortFieldMissing
	values  [][]byte
	tmp     [][]byte
}

func (s *SortField) UpdateVisitor(field string, term []byte) { _ = "STUB: not implemented"; return }

func (s *SortField) Value(i *DocumentMatch) string { _ = "STUB: not implemented"; return "" }

func (s *SortField) DecodeValue(value string) string { _ = "STUB: not implemented"; return "" }

func (s *SortField) Descending() bool { _ = "STUB: not implemented"; return false }

func (s *SortField) filterTermsByMode(terms [][]byte) string { _ = "STUB: not implemented"; return "" }

func (s *SortField) filterTermsByType(terms [][]byte) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func (s *SortField) RequiresDocID() bool { _ = "STUB: not implemented"; return false }

func (s *SortField) RequiresScoring() bool { _ = "STUB: not implemented"; return false }

func (s *SortField) RequiresFields() []string { _ = "STUB: not implemented"; return nil }

func (s *SortField) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SortField) Copy() SearchSort { _ = "STUB: not implemented"; return *new(SearchSort) }

func (s *SortField) Reverse() { _ = "STUB: not implemented"; return }

type SortDocID struct {
	Desc bool
}

func (s *SortDocID) UpdateVisitor(field string, term []byte) { _ = "STUB: not implemented"; return }

func (s *SortDocID) Value(i *DocumentMatch) string { _ = "STUB: not implemented"; return "" }

func (s *SortDocID) DecodeValue(value string) string { _ = "STUB: not implemented"; return "" }

func (s *SortDocID) Descending() bool { _ = "STUB: not implemented"; return false }

func (s *SortDocID) RequiresDocID() bool { _ = "STUB: not implemented"; return false }

func (s *SortDocID) RequiresScoring() bool { _ = "STUB: not implemented"; return false }

func (s *SortDocID) RequiresFields() []string { _ = "STUB: not implemented"; return nil }

func (s *SortDocID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SortDocID) Copy() SearchSort { _ = "STUB: not implemented"; return *new(SearchSort) }

func (s *SortDocID) Reverse() { _ = "STUB: not implemented"; return }

type SortScore struct {
	Desc bool
}

func (s *SortScore) UpdateVisitor(field string, term []byte) { _ = "STUB: not implemented"; return }

func (s *SortScore) Value(i *DocumentMatch) string { _ = "STUB: not implemented"; return "" }

func (s *SortScore) DecodeValue(value string) string { _ = "STUB: not implemented"; return "" }

func (s *SortScore) Descending() bool { _ = "STUB: not implemented"; return false }

func (s *SortScore) RequiresDocID() bool { _ = "STUB: not implemented"; return false }

func (s *SortScore) RequiresScoring() bool { _ = "STUB: not implemented"; return false }

func (s *SortScore) RequiresFields() []string { _ = "STUB: not implemented"; return nil }

func (s *SortScore) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SortScore) Copy() SearchSort { _ = "STUB: not implemented"; return *new(SearchSort) }

func (s *SortScore) Reverse() { _ = "STUB: not implemented"; return }

var maxDistance = string(numeric.MustNewPrefixCodedInt64(math.MaxInt64, 0))

func NewSortGeoDistance(field, unit string, lon, lat float64, desc bool) (
	*SortGeoDistance, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SortGeoDistance struct {
	Field    string
	Desc     bool
	Unit     string
	values   [][]byte
	Lon      float64
	Lat      float64
	unitMult float64
	tmp      []byte
}

func (s *SortGeoDistance) UpdateVisitor(field string, term []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *SortGeoDistance) Value(i *DocumentMatch) string { _ = "STUB: not implemented"; return "" }

func (s *SortGeoDistance) DecodeValue(value string) string { _ = "STUB: not implemented"; return "" }

func (s *SortGeoDistance) Descending() bool { _ = "STUB: not implemented"; return false }

func (s *SortGeoDistance) findPrefixCodedNumericTerm(terms [][]byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (s *SortGeoDistance) RequiresDocID() bool { _ = "STUB: not implemented"; return false }

func (s *SortGeoDistance) RequiresScoring() bool { _ = "STUB: not implemented"; return false }

func (s *SortGeoDistance) RequiresFields() []string { _ = "STUB: not implemented"; return nil }

func (s *SortGeoDistance) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SortGeoDistance) Copy() SearchSort { _ = "STUB: not implemented"; return *new(SearchSort) }

func (s *SortGeoDistance) Reverse() { _ = "STUB: not implemented"; return }

type BytesSlice [][]byte

func (p BytesSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (p BytesSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (p BytesSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
