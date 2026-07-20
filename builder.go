package bleve

import (
	"github.com/blevesearch/bleve/v2/mapping"
	index "github.com/blevesearch/bleve_index_api"
)

type builderImpl struct {
	b index.IndexBuilder
	m mapping.IndexMapping
}

func (b *builderImpl) Index(id string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *builderImpl) Close() error { _ = "STUB: not implemented"; return nil }

func newBuilder(path string, mapping mapping.IndexMapping, config map[string]interface{}) (Builder, error) {
	_ = "STUB: not implemented"
	return *new(Builder), nil
}
