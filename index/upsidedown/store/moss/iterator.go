package moss

import (
	"github.com/couchbase/moss"
)

type Iterator struct {
	store *Store
	ss    moss.Snapshot
	iter  moss.Iterator
	start []byte
	end   []byte
	k     []byte
	v     []byte
	err   error
}

func (x *Iterator) Seek(seekToKey []byte) { _ = "STUB: not implemented"; return }

func (x *Iterator) Next() { _ = "STUB: not implemented"; return }

func (x *Iterator) Current() ([]byte, []byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (x *Iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (x *Iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (x *Iterator) Valid() bool { _ = "STUB: not implemented"; return false }

func (x *Iterator) Close() error { _ = "STUB: not implemented"; return nil }

func (x *Iterator) current() { _ = "STUB: not implemented"; return }
