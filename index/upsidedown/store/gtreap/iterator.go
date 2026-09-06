package gtreap

import (
	"sync"

	"github.com/blevesearch/gtreap"
)

type Iterator struct {
	t *gtreap.Treap

	m        sync.Mutex
	cancelCh chan struct{}
	nextCh   chan *Item
	curr     *Item
	currOk   bool

	prefix []byte
	start  []byte
	end    []byte
}

func (w *Iterator) Seek(k []byte) { _ = "STUB: not implemented"; return }

func (w *Iterator) restart(start *Item) *Iterator { _ = "STUB: not implemented"; return nil }

func (w *Iterator) Next() { _ = "STUB: not implemented"; return }

func (w *Iterator) Current() ([]byte, []byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (w *Iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (w *Iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (w *Iterator) Valid() bool { _ = "STUB: not implemented"; return false }

func (w *Iterator) Close() error { _ = "STUB: not implemented"; return nil }
