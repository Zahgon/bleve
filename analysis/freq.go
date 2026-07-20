package analysis

import (
	index "github.com/blevesearch/bleve_index_api"
)

func TokenFrequency(tokens TokenStream, arrayPositions []uint64, options index.FieldIndexingOptions) index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}
