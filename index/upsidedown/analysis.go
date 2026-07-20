package upsidedown

import (
	index "github.com/blevesearch/bleve_index_api"
)

type IndexRow interface {
	KeySize() int
	KeyTo([]byte) (int, error)
	Key() []byte

	ValueSize() int
	ValueTo([]byte) (int, error)
	Value() []byte
}

type AnalysisResult struct {
	DocID string
	Rows  []IndexRow
}

func (udc *UpsideDownCouch) Analyze(d index.Document) *AnalysisResult {
	_ = "STUB: not implemented"
	return nil
}

func (udc *UpsideDownCouch) analyze(d index.Document) *AnalysisResult {
	_ = "STUB: not implemented"
	return nil
}
