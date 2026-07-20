package scorch

import (
	"reflect"

	segment "github.com/blevesearch/scorch_segment_api/v2"
)

var reflectStaticSizeIndexSnapshotThesaurusTermReader int

func init() {
	var istr IndexSnapshotThesaurusTermReader
	reflectStaticSizeIndexSnapshotThesaurusTermReader = int(reflect.TypeOf(istr).Size())
}

type IndexSnapshotThesaurusTermReader struct {
	name          string
	snapshot      *IndexSnapshot
	thesauri      []segment.Thesaurus
	postings      []segment.SynonymsList
	iterators     []segment.SynonymsIterator
	segmentOffset int
}

func (i *IndexSnapshotThesaurusTermReader) Size() int { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotThesaurusTermReader) Next() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (i *IndexSnapshotThesaurusTermReader) Close() error { _ = "STUB: not implemented"; return nil }
