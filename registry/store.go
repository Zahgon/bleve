package registry

import (
	store "github.com/blevesearch/upsidedown_store_api"
)

func RegisterKVStore(name string, constructor KVStoreConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type KVStoreConstructor func(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error)
type KVStoreRegistry map[string]KVStoreConstructor

func KVStoreConstructorByName(name string) KVStoreConstructor {
	_ = "STUB: not implemented"
	return *new(KVStoreConstructor)
}

func KVStoreTypesAndInstances() ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }
