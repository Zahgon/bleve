package gtreap

import (
	"github.com/blevesearch/gtreap"
)

type Reader struct {
	t *gtreap.Treap
}

func (w *Reader) Get(k []byte) (v []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reader) MultiGet(keys [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Reader) PrefixIterator(k []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (w *Reader) RangeIterator(start, end []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (w *Reader) Close() error { _ = "STUB: not implemented"; return nil }
