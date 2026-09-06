package metrics

import store "github.com/blevesearch/upsidedown_store_api"

type Iterator struct {
	s *Store
	o store.KVIterator
}

func (i *Iterator) Seek(x []byte) { _ = "STUB: not implemented"; return }

func (i *Iterator) Next() { _ = "STUB: not implemented"; return }

func (i *Iterator) Current() ([]byte, []byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (i *Iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (i *Iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (i *Iterator) Valid() bool { _ = "STUB: not implemented"; return false }

func (i *Iterator) Close() error { _ = "STUB: not implemented"; return nil }
