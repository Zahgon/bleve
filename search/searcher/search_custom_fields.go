package searcher

import (
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func loadDocValuesOnHitWithTypes(hit *search.DocumentMatch, dvReader index.DocValueReader,
	r index.IndexReader, fieldTypes map[string]string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addFieldValueToMap(fields map[string]interface{}, name string, value interface{}) {
	_ = "STUB: not implemented"
	return
}

func decodeDocValueTerm(term []byte, fieldType string) interface{} {
	_ = "STUB: not implemented"
	return nil
}
