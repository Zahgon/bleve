package metrics

import store "github.com/blevesearch/upsidedown_store_api"

type Reader struct {
	s *Store
	o store.KVReader
}

func (r *Reader) Get(key []byte) (v []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reader) MultiGet(keys [][]byte) (vals [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) PrefixIterator(prefix []byte) (i store.KVIterator) {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *Reader) RangeIterator(start, end []byte) (i store.KVIterator) {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }
