package upsidedown

import (
	store "github.com/blevesearch/upsidedown_store_api"
)

func dumpPrefix(kvreader store.KVReader, rv chan interface{}, prefix []byte) {
	_ = "STUB: not implemented"
	return
}

func dumpRange(kvreader store.KVReader, rv chan interface{}, start, end []byte) {
	_ = "STUB: not implemented"
	return
}

func (i *IndexReader) DumpAll() chan interface{} { _ = "STUB: not implemented"; return nil }

func (i *IndexReader) DumpFields() chan interface{} { _ = "STUB: not implemented"; return nil }

type keyset [][]byte

func (k keyset) Len() int           { _ = "STUB: not implemented"; return 0 }
func (k keyset) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (k keyset) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (i *IndexReader) DumpDoc(id string) chan interface{} { _ = "STUB: not implemented"; return nil }
