package scorch

import (
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type segmentThesCursor struct {
	thes segment.Thesaurus
	itr  segment.ThesaurusIterator
	curr index.ThesaurusEntry
}

type IndexSnapshotThesaurusKeys struct {
	snapshot *IndexSnapshot
	cursors  []*segmentThesCursor
	entry    index.ThesaurusEntry
}

func (i *IndexSnapshotThesaurusKeys) Len() int           { _ = "STUB: not implemented"; return 0 }
func (i *IndexSnapshotThesaurusKeys) Less(a, b int) bool { _ = "STUB: not implemented"; return false }

func (i *IndexSnapshotThesaurusKeys) Swap(a, b int) { _ = "STUB: not implemented"; return }

func (i *IndexSnapshotThesaurusKeys) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (i *IndexSnapshotThesaurusKeys) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshotThesaurusKeys) Next() (*index.ThesaurusEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshotThesaurusKeys) Close() error { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshotThesaurusKeys) Contains(key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
