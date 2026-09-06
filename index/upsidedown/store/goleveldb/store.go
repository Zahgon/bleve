package goleveldb

import (
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/goleveldb/leveldb"
	"github.com/blevesearch/goleveldb/leveldb/opt"
	store "github.com/blevesearch/upsidedown_store_api"
)

const (
	Name                    = "goleveldb"
	defaultCompactBatchSize = 250
)

type Store struct {
	path string
	opts *opt.Options
	db   *leveldb.DB
	mo   store.MergeOperator

	defaultWriteOptions *opt.WriteOptions
	defaultReadOptions  *opt.ReadOptions
}

func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	_ = "STUB: not implemented"
	return *new(store.KVStore), nil
}

func (ldbs *Store) Close() error { _ = "STUB: not implemented"; return nil }

func (ldbs *Store) Reader() (store.KVReader, error) {
	_ = "STUB: not implemented"
	return *new(store.KVReader), nil
}

func (ldbs *Store) Writer() (store.KVWriter, error) {
	_ = "STUB: not implemented"
	return *new(store.KVWriter), nil
}

func (ldbs *Store) CompactWithBatchSize(batchSize int) error { _ = "STUB: not implemented"; return nil }

func (ldbs *Store) Compact() error { _ = "STUB: not implemented"; return nil }

func init() {
	err := registry.RegisterKVStore(Name, New)
	if err != nil {
		panic(err)
	}
}
