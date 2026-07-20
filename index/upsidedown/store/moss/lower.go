package moss

import (
	"sync"

	"github.com/couchbase/moss"

	store "github.com/blevesearch/upsidedown_store_api"
)

func initLowerLevelStore(
	config map[string]interface{},
	lowerLevelStoreName string,
	lowerLevelStoreConfig map[string]interface{},
	lowerLevelMaxBatchSize uint64,
	options moss.CollectionOptions,
) (moss.Snapshot, moss.LowerLevelUpdate, store.KVStore, statsFunc, error) {
	_ = "STUB: not implemented"
	return *new(moss.Snapshot), *new(moss.LowerLevelUpdate), *new(store.KVStore), *new(statsFunc), nil
}

type llStore struct {
	kvStore store.KVStore

	config   map[string]interface{}
	llConfig map[string]interface{}

	logf func(format string, a ...interface{})

	m    sync.Mutex
	refs int
}

type llSnapshot struct {
	llStore        *llStore
	kvReader       store.KVReader
	childSnapshots map[string]*llSnapshot

	m    sync.Mutex
	refs int
}

type llIterator struct {
	llSnapshot *llSnapshot

	kvReader store.KVReader

	kvIterator store.KVIterator
}

type readerSource interface {
	Reader() (store.KVReader, error)
}

func (s *llStore) addRef() *llStore { _ = "STUB: not implemented"; return nil }

func (s *llStore) decRef() { _ = "STUB: not implemented"; return }

func (s *llStore) update(ssHigher moss.Snapshot, maxBatchSize uint64) (
	ssLower moss.Snapshot, err error,
) {
	_ = "STUB: not implemented"
	return *new(moss.Snapshot), nil
}

func (llss *llSnapshot) addRef() *llSnapshot { _ = "STUB: not implemented"; return nil }

func (llss *llSnapshot) decRef() { _ = "STUB: not implemented"; return }

func (llss *llSnapshot) ChildCollectionNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (llss *llSnapshot) ChildCollectionSnapshot(childCollectionName string) (
	moss.Snapshot, error,
) {
	_ = "STUB: not implemented"
	return *new(moss.Snapshot), nil
}

func (llss *llSnapshot) Close() error { _ = "STUB: not implemented"; return nil }

func (llss *llSnapshot) Get(key []byte,
	readOptions moss.ReadOptions,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (llss *llSnapshot) StartIterator(
	startKeyInclusive, endKeyExclusive []byte,
	iteratorOptions moss.IteratorOptions,
) (moss.Iterator, error) {
	_ = "STUB: not implemented"
	return *new(moss.Iterator), nil
}

func (lli *llIterator) Close() error { _ = "STUB: not implemented"; return nil }

func (lli *llIterator) Next() error { _ = "STUB: not implemented"; return nil }

func (lli *llIterator) SeekTo(k []byte) error { _ = "STUB: not implemented"; return nil }

func (lli *llIterator) Current() (key, val []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (lli *llIterator) CurrentEx() (
	entryEx moss.EntryEx, key, val []byte, err error,
) {
	_ = "STUB: not implemented"
	return *new(moss.EntryEx), nil, nil, nil
}

func InitMossStore(config map[string]interface{}, options moss.CollectionOptions) (
	moss.Snapshot, moss.LowerLevelUpdate, store.KVStore, statsFunc, error,
) {
	_ = "STUB: not implemented"
	return *new(moss.Snapshot), *new(moss.LowerLevelUpdate), *new(store.KVStore), *new(statsFunc), nil
}

type mossStoreWrapper struct {
	m    sync.Mutex
	refs int
	s    *moss.Store
}

func (w *mossStoreWrapper) AddRef() { _ = "STUB: not implemented"; return }

func (w *mossStoreWrapper) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (w *mossStoreWrapper) Reader() (store.KVReader, error) {
	_ = "STUB: not implemented"
	return *new(store.KVReader), nil
}

func (w *mossStoreWrapper) Writer() (store.KVWriter, error) {
	_ = "STUB: not implemented"
	return *new(store.KVWriter), nil
}

func (w *mossStoreWrapper) Actual() *moss.Store { _ = "STUB: not implemented"; return nil }

func (w *mossStoreWrapper) histograms() string { _ = "STUB: not implemented"; return "" }
