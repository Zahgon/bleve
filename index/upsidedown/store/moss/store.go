package moss

import (
	"encoding/json"
	"sync"

	"github.com/couchbase/moss"

	"github.com/blevesearch/bleve/v2/registry"
	store "github.com/blevesearch/upsidedown_store_api"
)

var RegistryCollectionOptions = map[string]moss.CollectionOptions{}

const Name = "moss"

type Store struct {
	m       sync.Mutex
	ms      moss.Collection
	mo      store.MergeOperator
	llstore store.KVStore
	llstats statsFunc

	s      *stats
	config map[string]interface{}
}

type statsFunc func() map[string]interface{}

func New(mo store.MergeOperator, config map[string]interface{}) (
	store.KVStore, error) {
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

func (s *Store) Logf(fmt string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *Store) Stats() json.Marshaler { _ = "STUB: not implemented"; return *new(json.Marshaler) }

func (s *Store) StatsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (s *Store) LowerLevelStore() store.KVStore {
	_ = "STUB: not implemented"
	return *new(store.KVStore)
}

func (s *Store) Collection() moss.Collection {
	_ = "STUB: not implemented"
	return *new(moss.Collection)
}

func init() {
	err := registry.RegisterKVStore(Name, New)
	if err != nil {
		panic(err)
	}
}
