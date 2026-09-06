package metrics

import (
	"container/list"
	"encoding/json"
	"io"
	"sync"

	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/go-metrics"
	store "github.com/blevesearch/upsidedown_store_api"
)

const Name = "metrics"

type Store struct {
	o store.KVStore

	timerReaderGet            metrics.Timer
	timerReaderMultiGet       metrics.Timer
	timerReaderPrefixIterator metrics.Timer
	timerReaderRangeIterator  metrics.Timer
	timerWriterExecuteBatch   metrics.Timer
	timerIteratorSeek         metrics.Timer
	timerIteratorNext         metrics.Timer
	timerBatchMerge           metrics.Timer

	m      sync.Mutex
	errors *list.List

	s *stats
}

func New(mo store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	_ = "STUB: not implemented"
	return *new(store.KVStore), nil
}

func init() {
	err := registry.RegisterKVStore(Name, New)
	if err != nil {
		panic(err)
	}
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

const MaxErrors = 100

type StoreError struct {
	Time string
	Op   string
	Err  string
	Key  string
}

func (s *Store) AddError(op string, err error, key []byte) { _ = "STUB: not implemented"; return }

func (s *Store) WriteJSON(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

func (s *Store) WriteCSVHeader(w io.Writer) { _ = "STUB: not implemented"; return }

func (s *Store) WriteCSV(w io.Writer) { _ = "STUB: not implemented"; return }

func (s *Store) Stats() json.Marshaler { _ = "STUB: not implemented"; return *new(json.Marshaler) }

func (s *Store) StatsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }
