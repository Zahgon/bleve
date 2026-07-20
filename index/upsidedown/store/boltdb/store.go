package boltdb

import (
	"encoding/json"

	"github.com/blevesearch/bleve/v2/registry"
	store "github.com/blevesearch/upsidedown_store_api"
	bolt "go.etcd.io/bbolt"
)

const (
	Name                    = "boltdb"
	defaultCompactBatchSize = 100
)

type Store struct {
	path        string
	bucket      string
	db          *bolt.DB
	noSync      bool
	fillPercent float64
	mo          store.MergeOperator
}

func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	_ = "STUB: not implemented"
	return *new(store.KVStore), nil
}

func (bs *Store) Close() error { _ = "STUB: not implemented"; return nil }

func (bs *Store) Reader() (store.KVReader, error) {
	_ = "STUB: not implemented"
	return *new(store.KVReader), nil
}

func (bs *Store) Writer() (store.KVWriter, error) {
	_ = "STUB: not implemented"
	return *new(store.KVWriter), nil
}

func (bs *Store) Stats() json.Marshaler { _ = "STUB: not implemented"; return *new(json.Marshaler) }

func (bs *Store) CompactWithBatchSize(batchSize int) error { _ = "STUB: not implemented"; return nil }

func (bs *Store) Compact() error { _ = "STUB: not implemented"; return nil }

func init() {
	err := registry.RegisterKVStore(Name, New)
	if err != nil {
		panic(err)
	}
}
