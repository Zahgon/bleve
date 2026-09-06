package scorch

import (
	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

type termDictionaryOmitCount interface {
	AutomatonIteratorOmitCount(a segment.Automaton,
		startKeyInclusive, endKeyExclusive []byte) segment.DictionaryIterator
}

func automatonIteratorOmitCount(dict segment.TermDictionary, a segment.Automaton,
	startKeyInclusive, endKeyExclusive []byte) segment.DictionaryIterator {
	_ = "STUB: not implemented"
	return *new(segment.DictionaryIterator)
}

type segmentDictCursor struct {
	dict segment.TermDictionary
	itr  segment.DictionaryIterator
	curr index.DictEntry
}

type IndexSnapshotFieldDict struct {
	cardinality int
	bytesRead   uint64

	snapshot *IndexSnapshot
	cursors  []*segmentDictCursor
	entry    index.DictEntry
}

func (i *IndexSnapshotFieldDict) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotFieldDict) Len() int           { _ = "STUB: not implemented"; return 0 }
func (i *IndexSnapshotFieldDict) Less(a, b int) bool { _ = "STUB: not implemented"; return false }

func (i *IndexSnapshotFieldDict) Swap(a, b int) { _ = "STUB: not implemented"; return }

func (i *IndexSnapshotFieldDict) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (i *IndexSnapshotFieldDict) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshotFieldDict) Next() (*index.DictEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshotFieldDict) Cardinality() int { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotFieldDict) Close() error { _ = "STUB: not implemented"; return nil }

func (i *IndexSnapshotFieldDict) Contains(key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
