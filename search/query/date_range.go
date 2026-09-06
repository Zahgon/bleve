package query

import (
	"context"
	"time"

	"github.com/blevesearch/bleve/v2/analysis/datetime/optional"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

var QueryDateTimeParser = optional.Name

var QueryDateTimeFormat = time.RFC3339

var cache = registry.NewCache()

type BleveQueryTime struct {
	time.Time
}

var MinRFC3339CompatibleTime time.Time
var MaxRFC3339CompatibleTime time.Time

func init() {
	MinRFC3339CompatibleTime, _ = time.Parse(time.RFC3339, "1677-12-01T00:00:00Z")
	MaxRFC3339CompatibleTime, _ = time.Parse(time.RFC3339, "2262-04-11T11:59:59Z")
}

func queryTimeFromString(t string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (t *BleveQueryTime) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *BleveQueryTime) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type DateRangeQuery struct {
	Start          BleveQueryTime `json:"start,omitempty"`
	End            BleveQueryTime `json:"end,omitempty"`
	InclusiveStart *bool          `json:"inclusive_start,omitempty"`
	InclusiveEnd   *bool          `json:"inclusive_end,omitempty"`
	FieldVal       string         `json:"field,omitempty"`
	BoostVal       *Boost         `json:"boost,omitempty"`
}

func NewDateRangeQuery(start, end time.Time) *DateRangeQuery { _ = "STUB: not implemented"; return nil }

func NewDateRangeInclusiveQuery(start, end time.Time, startInclusive, endInclusive *bool) *DateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *DateRangeQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *DateRangeQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *DateRangeQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *DateRangeQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *DateRangeQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *DateRangeQuery) parseEndpoints() (*float64, *float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (q *DateRangeQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func isDatetimeCompatible(t BleveQueryTime) bool { _ = "STUB: not implemented"; return false }
