package scorch

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/RoaringBitmap/roaring/v2"
	geov2 "github.com/blevesearch/bleve/v2/geov2"
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
	"github.com/blevesearch/vellum"
	lev "github.com/blevesearch/vellum/levenshtein"
)

var lb1, lb2 *lev.LevenshteinAutomatonBuilder

type asynchSegmentResult struct {
	dict    segment.TermDictionary
	dictItr segment.DictionaryIterator

	index int
	docs  *roaring.Bitmap

	thesItr segment.ThesaurusIterator

	err error
}

var reflectStaticSizeIndexSnapshot int
var reflectStaticSizeIndexSnapshotGeoShapeV2Reader int
var reflectStaticSizeRoaringIntIterator int

func init() {
	var is interface{} = IndexSnapshot{}
	reflectStaticSizeIndexSnapshot = int(reflect.TypeOf(is).Size())
	var err error
	lb1, err = lev.NewLevenshteinAutomatonBuilder(1, true)
	if err != nil {
		panic(fmt.Errorf("levenshtein automaton ed1 builder err: %v", err))
	}
	lb2, err = lev.NewLevenshteinAutomatonBuilder(2, true)
	if err != nil {
		panic(fmt.Errorf("levenshtein automaton ed2 builder err: %v", err))
	}
	var gcr IndexSnapshotGeoShapeV2Reader
	reflectStaticSizeIndexSnapshotGeoShapeV2Reader = int(reflect.TypeOf(gcr).Size())
	var rip roaring.IntIterator
	reflectStaticSizeRoaringIntIterator = int(reflect.TypeOf(rip).Size())
}

type IndexSnapshot struct {
	parent   *Scorch
	segment  []*SegmentSnapshot
	offsets  []uint64
	internal map[string][]byte
	epoch    uint64
	size     uint64
	creator  string

	m    sync.Mutex
	refs int64

	m2        sync.Mutex
	fieldTFRs map[string][]*IndexSnapshotTermFieldReader

	m3               sync.RWMutex
	fieldCardinality map[string]int

	updatedFields map[string]*index.UpdateFieldInfo

	fileWriterID string
}

func (i *IndexSnapshot) Segments() []*SegmentSnapshot { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshot) Internal() map[string][]byte { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshot) AddRef() { _ = "STUB: not implemented"; return }

func (i *IndexSnapshot) DecRef() (err error) { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshot) Close() error { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshot) Size() int { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshot) updateSize() { _ = "STUB: not implemented"; return }

func (is *IndexSnapshot) newIndexSnapshotFieldDict(field string,
	makeItr func(i segment.TermDictionary) segment.DictionaryIterator,
	randomLookup bool,
) (*IndexSnapshotFieldDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is *IndexSnapshot) FieldCardinality(field string) (rv int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (is *IndexSnapshot) FieldDict(field string) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func calculateExclusiveEndFromInclusiveEnd(inclusiveEnd []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (is *IndexSnapshot) FieldDictRange(field string, startTerm []byte,
	endTerm []byte,
) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func calculateExclusiveEndFromPrefix(in []byte) []byte { _ = "STUB: not implemented"; return nil }

func (is *IndexSnapshot) FieldDictPrefix(field string,
	termPrefix []byte,
) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (is *IndexSnapshot) FieldDictRegexp(field string,
	termRegex string,
) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (is *IndexSnapshot) FieldDictRegexpAutomaton(field string,
	termRegex string,
) (index.FieldDict, index.RegexAutomaton, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), *new(index.RegexAutomaton), nil
}

func (is *IndexSnapshot) fieldDictRegexp(field string,
	termRegex string,
) (index.FieldDict, index.RegexAutomaton, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), *new(index.RegexAutomaton), nil
}

func (is *IndexSnapshot) getLevAutomaton(term string,
	fuzziness uint8,
) (vellum.Automaton, error) {
	_ = "STUB: not implemented"
	return *new(vellum.Automaton), nil
}

func (is *IndexSnapshot) FieldDictFuzzy(field string,
	term string, fuzziness int, prefix string,
) (index.FieldDict, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), nil
}

func (is *IndexSnapshot) FieldDictFuzzyAutomaton(field string,
	term string, fuzziness int, prefix string,
) (index.FieldDict, index.FuzzyAutomaton, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), *new(index.FuzzyAutomaton), nil
}

func (is *IndexSnapshot) fieldDictFuzzy(field string,
	term string, fuzziness int, prefix string,
) (index.FieldDict, index.FuzzyAutomaton, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDict), *new(index.FuzzyAutomaton), nil
}

func (is *IndexSnapshot) FieldDictContains(field string) (index.FieldDictContains, error) {
	_ = "STUB: not implemented"
	return *new(index.FieldDictContains), nil
}

func (is *IndexSnapshot) DocIDReaderAll() (index.DocIDReader, error) {
	_ = "STUB: not implemented"
	return *new(index.DocIDReader), nil
}

func (is *IndexSnapshot) DocIDReaderOnly(ids []string) (index.DocIDReader, error) {
	_ = "STUB: not implemented"
	return *new(index.DocIDReader), nil
}

func (is *IndexSnapshot) newDocIDReader(results chan *asynchSegmentResult) (index.DocIDReader, error) {
	_ = "STUB: not implemented"
	return *new(index.DocIDReader), nil
}

func (is *IndexSnapshot) Fields() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (is *IndexSnapshot) GetInternal(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is *IndexSnapshot) DocCount() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (is *IndexSnapshot) Document(id string) (rv index.Document, err error) {
	_ = "STUB: not implemented"
	return *new(index.Document), nil
}

func (is *IndexSnapshot) segmentIndexAndLocalDocNumFromGlobal(docNum uint64) (int, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (is *IndexSnapshot) ExternalID(id index.IndexInternalID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (is *IndexSnapshot) segmentIndexAndLocalDocNum(id index.IndexInternalID) (int, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (is *IndexSnapshot) InternalID(id string) (rv index.IndexInternalID, err error) {
	_ = "STUB: not implemented"
	return *new(index.IndexInternalID), nil
}

func (is *IndexSnapshot) TermFieldReader(ctx context.Context, term []byte, field string, includeFreq,
	includeNorm, includeTermVectors bool,
) (index.TermFieldReader, error) {
	_ = "STUB: not implemented"
	return *new(index.TermFieldReader), nil
}

func (is *IndexSnapshot) allocTermFieldReaderDicts(field string) (tfr *IndexSnapshotTermFieldReader) {
	_ = "STUB: not implemented"
	return nil
}

var DefaultFieldTFRCacheThreshold int = 0

func (is *IndexSnapshot) getFieldTFRCacheThreshold() int { _ = "STUB: not implemented"; return 0 }

func (is *IndexSnapshot) recycleTermFieldReader(tfr *IndexSnapshotTermFieldReader) {
	_ = "STUB: not implemented"
	return
}

func (is *IndexSnapshot) documentVisitFieldTermsOnSegment(
	segmentIndex int, localDocNum uint64, fields []string, cFields []string,
	visitor index.DocValueVisitor, dvs segment.DocVisitState) (
	cFieldsOut []string, dvsOut segment.DocVisitState, err error,
) {
	_ = "STUB: not implemented"
	return nil, *new(segment.DocVisitState), nil
}

func (is *IndexSnapshot) DocValueReader(fields []string) (
	index.DocValueReader, error,
) {
	_ = "STUB: not implemented"
	return *new(index.DocValueReader), nil
}

type DocValueReader struct {
	i      *IndexSnapshot
	fields []string
	dvs    segment.DocVisitState

	currSegmentIndex int
	currCachedFields []string

	totalBytesRead uint64
	bytesRead      uint64
}

func (dvr *DocValueReader) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (dvr *DocValueReader) VisitDocValues(id index.IndexInternalID,
	visitor index.DocValueVisitor,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (is *IndexSnapshot) DumpAll() chan interface{} { _ = "STUB: not implemented"; return nil }

func (is *IndexSnapshot) DumpDoc(id string) chan interface{} { _ = "STUB: not implemented"; return nil }

func (is *IndexSnapshot) DumpFields() chan interface{} { _ = "STUB: not implemented"; return nil }

func (is *IndexSnapshot) diskSegmentsPaths() map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (is *IndexSnapshot) reClaimableDocsRatio() float64 { _ = "STUB: not implemented"; return 0 }

func subtractStrings(a, b []string) []string { _ = "STUB: not implemented"; return nil }

func (is *IndexSnapshot) CopyTo(d index.Directory) error { _ = "STUB: not implemented"; return nil }

func (is *IndexSnapshot) UpdateIOStats(val uint64) { _ = "STUB: not implemented"; return }

func (is *IndexSnapshot) GetSpatialAnalyzerPlugin(typ string) (
	index.SpatialAnalyzerPlugin, error,
) {
	_ = "STUB: not implemented"
	return *new(index.SpatialAnalyzerPlugin), nil
}

func (is *IndexSnapshot) CloseCopyReader() error { _ = "STUB: not implemented"; return nil }

func (is *IndexSnapshot) ThesaurusTermReader(ctx context.Context, thesaurusName string, term []byte) (index.ThesaurusTermReader, error) {
	_ = "STUB: not implemented"
	return *new(index.ThesaurusTermReader), nil
}

func (is *IndexSnapshot) newIndexSnapshotThesaurusKeys(name string,
	makeItr func(i segment.Thesaurus) segment.ThesaurusIterator,
) (*IndexSnapshotThesaurusKeys, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is *IndexSnapshot) ThesaurusKeys(name string) (index.ThesaurusKeys, error) {
	_ = "STUB: not implemented"
	return *new(index.ThesaurusKeys), nil
}

func (is *IndexSnapshot) ThesaurusKeysFuzzy(name string,
	term string, fuzziness int, prefix string,
) (index.ThesaurusKeys, error) {
	_ = "STUB: not implemented"
	return *new(index.ThesaurusKeys), nil
}

func (is *IndexSnapshot) ThesaurusKeysPrefix(name string,
	termPrefix []byte,
) (index.ThesaurusKeys, error) {
	_ = "STUB: not implemented"
	return *new(index.ThesaurusKeys), nil
}

func (is *IndexSnapshot) ThesaurusKeysRegexp(name string,
	termRegex string,
) (index.ThesaurusKeys, error) {
	_ = "STUB: not implemented"
	return *new(index.ThesaurusKeys), nil
}

func (is *IndexSnapshot) UpdateSynonymSearchCount(delta uint64) { _ = "STUB: not implemented"; return }

func (is *IndexSnapshot) UpdateFieldsInfo(updatedFields map[string]*index.UpdateFieldInfo) {
	_ = "STUB: not implemented"
	return
}

func (is *IndexSnapshot) MergeUpdateFieldsInfo(updatedFields map[string]*index.UpdateFieldInfo) {
	_ = "STUB: not implemented"
	return
}

func (is *IndexSnapshot) TermFrequencies(field string, limit int, descending bool) (
	termFreqs []index.TermFreq, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshot) Ancestors(ID index.IndexInternalID, prealloc []index.AncestorID) ([]index.AncestorID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshot) GeoShapeV2FieldReader(ctx context.Context, field string) (
	index.GeoShapeV2FieldReader, error) {
	_ = "STUB: not implemented"
	return *new(index.GeoShapeV2FieldReader), nil
}

type IndexSnapshotGeoShapeV2Reader struct {
	field string

	postings      []*roaring.Bitmap
	iterators     []roaring.IntPeekable
	segmentOffset int

	snapshot *IndexSnapshot
}

func (g *IndexSnapshotGeoShapeV2Reader) Search(shape index.GeoJSON,
	relation string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *IndexSnapshotGeoShapeV2Reader) searchSeg(segID int,
	query geov2.Query) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *IndexSnapshotGeoShapeV2Reader) Next(preAlloced *index.GeoShapeV2FieldDoc) (
	*index.GeoShapeV2FieldDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *IndexSnapshotGeoShapeV2Reader) Advance(ID index.IndexInternalID,
	preAlloced *index.GeoShapeV2FieldDoc) (*index.GeoShapeV2FieldDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *IndexSnapshotGeoShapeV2Reader) Close() error { _ = "STUB: not implemented"; return nil }

func (g *IndexSnapshotGeoShapeV2Reader) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (g *IndexSnapshotGeoShapeV2Reader) Size() int { _ = "STUB: not implemented"; return 0 }
