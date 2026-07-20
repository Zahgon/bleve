package document

import (
	index "github.com/blevesearch/bleve_index_api"
)

type Field interface {
	Name() string

	ArrayPositions() []uint64
	Options() index.FieldIndexingOptions
	Analyze()
	Value() []byte

	NumPlainTextBytes() uint64

	Size() int

	EncodedFieldType() byte
	AnalyzedLength() int
	AnalyzedTokenFrequencies() index.TokenFrequencies
}
