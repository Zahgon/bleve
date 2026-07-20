package boltdb

import (
	bolt "go.etcd.io/bbolt"
)

type Iterator struct {
	store  *Store
	tx     *bolt.Tx
	cursor *bolt.Cursor
	prefix []byte
	start  []byte
	end    []byte
	valid  bool
	key    []byte
	val    []byte
}

func (i *Iterator) updateValid() { _ = "STUB: not implemented"; return }

func (i *Iterator) Seek(k []byte) { _ = "STUB: not implemented"; return }

func (i *Iterator) Next() { _ = "STUB: not implemented"; return }

func (i *Iterator) Current() ([]byte, []byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (i *Iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (i *Iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (i *Iterator) Valid() bool { _ = "STUB: not implemented"; return false }

func (i *Iterator) Close() error { _ = "STUB: not implemented"; return nil }
