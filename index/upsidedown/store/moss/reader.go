package moss

import (
	"github.com/couchbase/moss"

	store "github.com/blevesearch/upsidedown_store_api"
)

type Reader struct {
	store *Store
	ss    moss.Snapshot
}

func (r *Reader) Get(k []byte) (v []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reader) MultiGet(keys [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) PrefixIterator(k []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *Reader) RangeIterator(start, end []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }

func incrementBytes(in []byte) []byte { _ = "STUB: not implemented"; return nil }
