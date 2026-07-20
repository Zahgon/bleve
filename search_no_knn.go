//go:build !vectors
// +build !vectors

package bleve

import (
	"context"
	"sort"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/collector"
	"github.com/blevesearch/bleve/v2/search/query"
	index "github.com/blevesearch/bleve_index_api"
)

const supportForVectorSearch = false

type SearchRequest struct {
	ClientContextID  string            `json:"client_context_id,omitempty"`
	Query            query.Query       `json:"query"`
	Size             int               `json:"size"`
	From             int               `json:"from"`
	Highlight        *HighlightRequest `json:"highlight,omitempty"`
	Fields           []string          `json:"fields,omitempty"`
	Facets           FacetsRequest     `json:"facets,omitempty"`
	Explain          bool              `json:"explain"`
	Sort             search.SortOrder  `json:"sort"`
	IncludeLocations bool              `json:"includeLocations"`
	Score            string            `json:"score,omitempty"`
	SearchAfter      []string          `json:"search_after,omitempty"`
	SearchBefore     []string          `json:"search_before,omitempty"`

	PreSearchData map[string]interface{} `json:"pre_search_data,omitempty"`

	Params *RequestParams `json:"params,omitempty"`

	sortFunc func(sort.Interface)
}

func (r *SearchRequest) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func copySearchRequest(req *SearchRequest, preSearchData map[string]interface{}) *SearchRequest {
	_ = "STUB: not implemented"
	return nil
}

func validateKNN(req *SearchRequest) error { _ = "STUB: not implemented"; return nil }

func (i *indexImpl) runKnnCollector(ctx context.Context, req *SearchRequest, reader index.IndexReader, preSearch bool) ([]*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setKnnHitsInCollector(knnHits []*search.DocumentMatch, coll *collector.TopNCollector) {
	_ = "STUB: not implemented"
	return
}

func requestHasKNN(req *SearchRequest) bool { _ = "STUB: not implemented"; return false }

func numKNNQueries(req *SearchRequest) int { _ = "STUB: not implemented"; return 0 }

func addKnnToDummyRequest(dummyReq *SearchRequest, realReq *SearchRequest) {
	_ = "STUB: not implemented"
	return
}

func validateAndDistributeKNNHits(knnHits []*search.DocumentMatch, indexes []Index) (map[string][]*search.DocumentMatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isKNNrequestSatisfiedByPreSearch(req *SearchRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func constructKnnPreSearchData(mergedOut map[string]map[string]interface{}, preSearchResult *SearchResult,
	indexes []Index) (map[string]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finalizeKNNResults(req *SearchRequest, knnHits []*search.DocumentMatch) []*search.DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func newKnnPreSearchResultProcessor(req *SearchRequest) *knnPreSearchResultProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (r *rescorer) prepareKnnRequest() { _ = "STUB: not implemented"; return }

func (r *rescorer) restoreKnnRequest() { _ = "STUB: not implemented"; return }
