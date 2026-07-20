package registry

import (
	index "github.com/blevesearch/bleve_index_api"
)

func RegisterIndexType(name string, constructor IndexTypeConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type IndexTypeConstructor func(storeName string, storeConfig map[string]interface{}, analysisQueue *index.AnalysisQueue) (index.Index, error)
type IndexTypeRegistry map[string]IndexTypeConstructor

func IndexTypeConstructorByName(name string) IndexTypeConstructor {
	_ = "STUB: not implemented"
	return *new(IndexTypeConstructor)
}

func IndexTypesAndInstances() ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }
