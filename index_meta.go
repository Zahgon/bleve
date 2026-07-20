package bleve

import (
	"github.com/blevesearch/bleve/v2/util"
	index "github.com/blevesearch/bleve_index_api"
)

const metaFilename = "index_meta.json"

type indexMeta struct {
	Storage    string                 `json:"storage"`
	IndexType  string                 `json:"index_type"`
	Config     map[string]interface{} `json:"config,omitempty"`
	fileWriter util.FileWriter
	fileReader util.FileReader
}

func newIndexMeta(indexType string, storage string, config map[string]interface{}, path string) (*indexMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openIndexMeta(path string) (*indexMeta, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *indexMeta) Save(path string) (err error) { _ = "STUB: not implemented"; return nil }

func (i *indexMeta) CopyTo(path string, d index.Directory) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (i *indexMeta) UpdateWriter(path string) error { _ = "STUB: not implemented"; return nil }

func indexMetaPath(path string) string { _ = "STUB: not implemented"; return "" }
