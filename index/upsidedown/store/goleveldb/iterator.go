package goleveldb

import "github.com/blevesearch/goleveldb/leveldb/iterator"

type Iterator struct {
	store    *Store
	iterator iterator.Iterator
}

func (ldi *Iterator) Seek(key []byte) { _ = "STUB: not implemented"; return }

func (ldi *Iterator) Next() { _ = "STUB: not implemented"; return }

func (ldi *Iterator) Current() ([]byte, []byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (ldi *Iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (ldi *Iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (ldi *Iterator) Valid() bool { _ = "STUB: not implemented"; return false }

func (ldi *Iterator) Close() error { _ = "STUB: not implemented"; return nil }
