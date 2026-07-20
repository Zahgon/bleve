package upsidedown

import (
	index "github.com/blevesearch/bleve_index_api"
	store "github.com/blevesearch/upsidedown_store_api"
)

type UpsideDownCouchFieldDict struct {
	indexReader *IndexReader
	iterator    store.KVIterator
	dictRow     *DictionaryRow
	dictEntry   *index.DictEntry
	field       uint16
}

func newUpsideDownCouchFieldDict(indexReader *IndexReader, field uint16, startTerm, endTerm []byte) (*UpsideDownCouchFieldDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpsideDownCouchFieldDict) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *UpsideDownCouchFieldDict) Next() (*index.DictEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpsideDownCouchFieldDict) Cardinality() int { _ = "STUB: not implemented"; return 0 }

func (r *UpsideDownCouchFieldDict) Close() error { _ = "STUB: not implemented"; return nil }
