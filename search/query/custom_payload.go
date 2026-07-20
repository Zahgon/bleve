package query

import (
	"github.com/blevesearch/bleve/v2/mapping"
	index "github.com/blevesearch/bleve_index_api"
)

func unmarshalCustomQueryPayload(data []byte, key string) (Query, []string, map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return *new(Query), nil, nil, nil
}

func resolveFieldTypes(fields []string, m mapping.IndexMapping) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func expandFieldWildcard(fields []string, i index.IndexReader) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
