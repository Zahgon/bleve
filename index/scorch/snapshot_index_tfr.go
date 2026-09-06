package scorch

import (
	"context"
	"reflect"

	index "github.com/blevesearch/bleve_index_api"
	segment "github.com/blevesearch/scorch_segment_api/v2"
)

var reflectStaticSizeIndexSnapshotTermFieldReader int

func init() {
	var istfr IndexSnapshotTermFieldReader
	reflectStaticSizeIndexSnapshotTermFieldReader = int(reflect.TypeOf(istfr).Size())
}

type IndexSnapshotTermFieldReader struct {
	term               []byte
	field              string
	snapshot           *IndexSnapshot
	dicts              []segment.TermDictionary
	postings           []segment.PostingsList
	iterators          []segment.PostingsIterator
	segmentOffset      int
	includeFreq        bool
	includeNorm        bool
	includeTermVectors bool
	currPosting        segment.Posting
	currID             index.IndexInternalID
	recycle            bool
	bytesRead          uint64
	ctx                context.Context
	unadorned          bool

	updateBytesRead bool
}

func (i *IndexSnapshotTermFieldReader) incrementBytesRead(val uint64) {
	_ = "STUB: not implemented"
	return
}

func (i *IndexSnapshotTermFieldReader) Size() int { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotTermFieldReader) Next(preAlloced *index.TermFieldDoc) (*index.TermFieldDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshotTermFieldReader) postingToTermFieldDoc(next segment.Posting, rv *index.TermFieldDoc) {
	_ = "STUB: not implemented"
	return
}

func (i *IndexSnapshotTermFieldReader) Advance(ID index.IndexInternalID, preAlloced *index.TermFieldDoc) (*index.TermFieldDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexSnapshotTermFieldReader) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (i *IndexSnapshotTermFieldReader) Close() error { _ = "STUB: not implemented"; return nil }
