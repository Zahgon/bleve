package null

import (
	"github.com/blevesearch/bleve/v2/registry"
	store "github.com/blevesearch/upsidedown_store_api"
)

const Name = "null"

type Store struct{}

func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	_ = "STUB: not implemented"
	return *new(store.KVStore), nil
}

func (i *Store) Close() error { _ = "STUB: not implemented"; return nil }

func (i *Store) Reader() (store.KVReader, error) {
	_ = "STUB: not implemented"
	return *new(store.KVReader), nil
}

func (i *Store) Writer() (store.KVWriter, error) {
	_ = "STUB: not implemented"
	return *new(store.KVWriter), nil
}

type reader struct{}

func (r *reader) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *reader) MultiGet(keys [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *reader) PrefixIterator(prefix []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *reader) RangeIterator(start, end []byte) store.KVIterator {
	_ = "STUB: not implemented"
	return *new(store.KVIterator)
}

func (r *reader) Close() error { _ = "STUB: not implemented"; return nil }

type iterator struct{}

func (i *iterator) SeekFirst()    { _ = "STUB: not implemented"; return }
func (i *iterator) Seek(k []byte) { _ = "STUB: not implemented"; return }
func (i *iterator) Next()         { _ = "STUB: not implemented"; return }

func (i *iterator) Current() ([]byte, []byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (i *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (i *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (i *iterator) Valid() bool { _ = "STUB: not implemented"; return false }

func (i *iterator) Close() error { _ = "STUB: not implemented"; return nil }

type batch struct{}

func (i *batch) Set(key, val []byte)   { _ = "STUB: not implemented"; return }
func (i *batch) Delete(key []byte)     { _ = "STUB: not implemented"; return }
func (i *batch) Merge(key, val []byte) { _ = "STUB: not implemented"; return }
func (i *batch) Reset()                { _ = "STUB: not implemented"; return }
func (i *batch) Close() error          { _ = "STUB: not implemented"; return nil }

type writer struct{}

func (w *writer) NewBatch() store.KVBatch { _ = "STUB: not implemented"; return *new(store.KVBatch) }

func (w *writer) NewBatchEx(options store.KVBatchOptions) ([]byte, store.KVBatch, error) {
	_ = "STUB: not implemented"
	return nil, *new(store.KVBatch), nil
}

func (w *writer) ExecuteBatch(store.KVBatch) error { _ = "STUB: not implemented"; return nil }

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func init() {
	err := registry.RegisterKVStore(Name, New)
	if err != nil {
		panic(err)
	}
}
