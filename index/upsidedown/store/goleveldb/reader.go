package goleveldb

import (
	"github.com/blevesearch/goleveldb/leveldb"
	store "github.com/blevesearch/upsidedown_store_api"
)

type Reader struct {
	store    *Store
	snapshot *leveldb.Snapshot
}

func (r *Reader) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reader) MultiGet(keys [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) PrefixIterator(prefix []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *Reader) RangeIterator(start, end []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }
