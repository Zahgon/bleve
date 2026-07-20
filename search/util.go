package search

import (
	"context"

	"github.com/blevesearch/geo/s2"
)

func MergeLocations(locations []FieldTermLocationMap) FieldTermLocationMap {
	_ = "STUB: not implemented"
	return *new(FieldTermLocationMap)
}

func MergeTermLocationMaps(rv, other TermLocationMap) TermLocationMap {
	_ = "STUB: not implemented"
	return *new(TermLocationMap)
}

func MergeFieldTermLocations(dest []FieldTermLocation, matches []*DocumentMatch) []FieldTermLocation {
	_ = "STUB: not implemented"
	return nil
}

func MergeFieldTermLocationsFromMatch(dest []FieldTermLocation, match *DocumentMatch) []FieldTermLocation {
	_ = "STUB: not implemented"
	return nil
}

func mergeFieldTermLocationFromMatch(dest []FieldTermLocation, dm *DocumentMatch) []FieldTermLocation {
	_ = "STUB: not implemented"
	return nil
}

type (
	SearchIncrementalCostCallbackMsg uint
	SearchQueryType                  uint
)

const (
	Term = SearchQueryType(1 << iota)
	Geo
	Numeric
	GenericCost
)

const (
	AddM = SearchIncrementalCostCallbackMsg(1 << iota)
	AbortM
	DoneM
)

type ContextKey string

func (c ContextKey) String() string { _ = "STUB: not implemented"; return "" }

const (
	SearchIncrementalCostKey ContextKey = "_search_incremental_cost_key"
	QueryTypeKey             ContextKey = "_query_type_key"
	FuzzyMatchPhraseKey      ContextKey = "_fuzzy_match_phrase_key"
	IncludeScoreBreakdownKey ContextKey = "_include_score_breakdown_key"

	PreSearchKey ContextKey = "_presearch_key"

	GetScoringModelCallbackKey ContextKey = "_get_scoring_model"

	SearchIOStatsCallbackKey ContextKey = "_search_io_stats_callback_key"

	GeoBufferPoolCallbackKey ContextKey = "_geo_buffer_pool_callback_key"

	SearchTypeKey ContextKey = "_search_type_key"

	SearcherStartCallbackKey ContextKey = "_searcher_start_callback_key"
	SearcherEndCallbackKey   ContextKey = "_searcher_end_callback_key"

	FieldTermSynonymMapKey ContextKey = "_field_term_synonym_map_key"

	BM25StatsKey ContextKey = "_bm25_stats_key"

	ScoreFusionKey ContextKey = "_fusion_rescoring_key"

	NestedSearchKey ContextKey = "_nested_search_key"
)

func RecordSearchCost(ctx context.Context,
	msg SearchIncrementalCostCallbackMsg, bytes uint64,
) {
	_ = "STUB: not implemented"
	return
}

const (
	MaxGeoBufPoolSize = 24 * 1024
	MinGeoBufPoolSize = 24
)

const (
	KnnPreSearchDataKey     = "_knn_pre_search_data_key"
	SynonymPreSearchDataKey = "_synonym_pre_search_data_key"
	BM25PreSearchDataKey    = "_bm25_pre_search_data_key"
)

const GlobalScoring = "_global_scoring"

type (
	SearcherStartCallbackFn func(size uint64) error

	SearcherEndCallbackFn func(size uint64) error

	GetScoringModelCallbackFn func() string

	HybridMergeCallbackFn func(ftsMatch *DocumentMatch, knnMatch *DocumentMatch)

	DescendantAdderCallbackFn func(parent *DocumentMatch, descendant *DocumentMatch) error

	GeoBufferPoolCallbackFunc func() *s2.GeoBufferPool

	SearchIOStatsCallbackFunc func(uint64)

	SearchIncrementalCostCallbackFn func(SearchIncrementalCostCallbackMsg,
		SearchQueryType, uint64)
)

type FieldTermSynonymMap map[string]map[string][]string

func (f FieldTermSynonymMap) MergeWith(fts FieldTermSynonymMap) { _ = "STUB: not implemented"; return }

var (
	BM25_k1 float64 = 1.2
	BM25_b  float64 = 0.75
)

type BM25Stats struct {
	DocCount         float64        `json:"doc_count"`
	FieldCardinality map[string]int `json:"field_cardinality"`
}

type FieldSet map[string]struct{}

func NewFieldSet() FieldSet { _ = "STUB: not implemented"; return *new(FieldSet) }

func (fs FieldSet) AddField(field string) { _ = "STUB: not implemented"; return }

func (fs FieldSet) HasID() bool { _ = "STUB: not implemented"; return false }

func (fs FieldSet) HasAll() bool { _ = "STUB: not implemented"; return false }
