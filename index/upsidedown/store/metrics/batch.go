package metrics

import store "github.com/blevesearch/upsidedown_store_api"

type Batch struct {
	s *Store
	o store.KVBatch
}

func (b *Batch) Set(key, val []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) Merge(key, val []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) Reset() { _ = "STUB: not implemented"; return }

func (b *Batch) Close() error { _ = "STUB: not implemented"; return nil }
