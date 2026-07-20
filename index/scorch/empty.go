package scorch

import segment "github.com/blevesearch/scorch_segment_api/v2"

type emptyPostingsIterator struct{}

func (e *emptyPostingsIterator) Next() (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (e *emptyPostingsIterator) Advance(uint64) (segment.Posting, error) {
	_ = "STUB: not implemented"
	return *new(segment.Posting), nil
}

func (e *emptyPostingsIterator) Size() int { _ = "STUB: not implemented"; return 0 }

func (e *emptyPostingsIterator) BytesRead() uint64 { _ = "STUB: not implemented"; return 0 }

func (e *emptyPostingsIterator) ResetBytesRead(uint64) { _ = "STUB: not implemented"; return }

func (e *emptyPostingsIterator) BytesWritten() uint64 { _ = "STUB: not implemented"; return 0 }

var anEmptyPostingsIterator = &emptyPostingsIterator{}
