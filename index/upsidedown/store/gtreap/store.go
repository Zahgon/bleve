package gtreap

import (
	"sync"

	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/gtreap"
	store "github.com/blevesearch/upsidedown_store_api"
)

const Name = "gtreap"

type Store struct {
	m  sync.Mutex
	t  *gtreap.Treap
	mo store.MergeOperator
}

type Item struct {
	k []byte
	v []byte
}

func itemCompare(a, b interface{}) int { _ = "STUB: not implemented"; return 0 }

func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	_ = "STUB: not implemented"
	return *new(store.KVStore), nil
}

func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Store) Reader() (store.KVReader, error) {
	_ = "STUB: not implemented"
	return *new(store.KVReader), nil
}

func (s *Store) Writer() (store.KVWriter, error) {
	_ = "STUB: not implemented"
	return *new(store.KVWriter), nil
}

func init() {
	err := registry.RegisterKVStore(Name, New)
	if err != nil {
		panic(err)
	}
}
