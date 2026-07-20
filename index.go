package bleve

import (
	"context"

	"github.com/blevesearch/bleve/v2/document"
	"github.com/blevesearch/bleve/v2/mapping"
	index "github.com/blevesearch/bleve_index_api"
)

type Batch struct {
	index    Index
	internal *index.Batch

	lastDocSize uint64
	totalSize   uint64
}

func (b *Batch) Index(id string, data interface{}) error { _ = "STUB: not implemented"; return nil }

func (b *Batch) IndexSynonym(id string, collection string, definition *SynonymDefinition) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Batch) LastDocSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *Batch) TotalDocsSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *Batch) IndexAdvanced(doc *document.Document) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Batch) Delete(id string) { _ = "STUB: not implemented"; return }

func (b *Batch) SetInternal(key, val []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) DeleteInternal(key []byte) { _ = "STUB: not implemented"; return }

func (b *Batch) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *Batch) String() string { _ = "STUB: not implemented"; return "" }

func (b *Batch) Reset() { _ = "STUB: not implemented"; return }

func (b *Batch) Merge(o *Batch) { _ = "STUB: not implemented"; return }

func (b *Batch) SetPersistedCallback(f index.BatchCallback) { _ = "STUB: not implemented"; return }

func (b *Batch) PersistedCallback() index.BatchCallback {
	_ = "STUB: not implemented"
	return *new(index.BatchCallback)
}

type Index interface {
	Index(id string, data interface{}) error
	Delete(id string) error

	NewBatch() *Batch
	Batch(b *Batch) error

	Document(id string) (index.Document, error)

	DocCount() (uint64, error)

	Search(req *SearchRequest) (*SearchResult, error)
	SearchInContext(ctx context.Context, req *SearchRequest) (*SearchResult, error)

	Fields() ([]string, error)

	FieldDict(field string) (index.FieldDict, error)
	FieldDictRange(field string, startTerm []byte, endTerm []byte) (index.FieldDict, error)
	FieldDictPrefix(field string, termPrefix []byte) (index.FieldDict, error)

	Close() error

	Mapping() mapping.IndexMapping

	Stats() *IndexStat
	StatsMap() map[string]interface{}

	GetInternal(key []byte) ([]byte, error)
	SetInternal(key, val []byte) error
	DeleteInternal(key []byte) error

	Name() string

	SetName(string)

	Advanced() (index.Index, error)
}

func New(path string, mapping mapping.IndexMapping) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func NewMemOnly(mapping mapping.IndexMapping) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func NewUsing(path string, mapping mapping.IndexMapping, indexType string, kvstore string, kvconfig map[string]interface{}) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func Open(path string) (Index, error) { _ = "STUB: not implemented"; return *new(Index), nil }

func OpenUsing(path string, runtimeConfig map[string]interface{}) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

type Builder interface {
	Index(id string, data interface{}) error
	Close() error
}

func NewBuilder(path string, mapping mapping.IndexMapping, config map[string]interface{}) (Builder, error) {
	_ = "STUB: not implemented"
	return *new(Builder), nil
}

type IndexCopyable interface {
	CopyTo(d index.Directory) error
}

type FileSystemDirectory string

type SynonymDefinition struct {
	Input []string `json:"input,omitempty"`

	Synonyms []string `json:"synonyms"`
}

func (sd *SynonymDefinition) Validate() error { _ = "STUB: not implemented"; return nil }

type SynonymIndex interface {
	Index

	IndexSynonym(id string, collection string, definition *SynonymDefinition) error
}

type IndexWithCallbacks interface {
	FileWriterIDsInUse() (map[string]struct{}, error)
	DropFileWriterIDs(ids map[string]struct{}) error
}

type InsightsIndex interface {
	Index

	TermFrequencies(field string, limit int, descending bool) ([]index.TermFreq, error)

	CentroidCardinalities(field string, limit int, desceding bool) ([]index.CentroidCardinality, error)
}

type TrainableIndex interface {
	Index
	Train(*Batch) error
}

type IndexFileCopyable interface {
	SetPathInBolt(key []byte, value []byte) error
	CopyFile(file string, d index.IndexDirectory) error
}
