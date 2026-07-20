package goleveldb

import (
	"github.com/blevesearch/goleveldb/leveldb"
	store "github.com/blevesearch/upsidedown_store_api"
)

type Batch struct {
	store *Store
	merge *store.EmulatedMerge
	batch *leveldb.Batch
}

func (b *Batch) Set(key, val []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) Delete(key []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) Merge(key, val []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) Reset() { _ = "STUB: not implemented"; return }

func (b *Batch) Close() error { _ = "STUB: not implemented"; return nil }
