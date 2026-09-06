package query

import (
	"context"
	"time"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

type DateRangeStringQuery struct {
	Start          string `json:"start,omitempty"`
	End            string `json:"end,omitempty"`
	InclusiveStart *bool  `json:"inclusive_start,omitempty"`
	InclusiveEnd   *bool  `json:"inclusive_end,omitempty"`
	FieldVal       string `json:"field,omitempty"`
	BoostVal       *Boost `json:"boost,omitempty"`
	DateTimeParser string `json:"datetime_parser,omitempty"`
}

func NewDateRangeStringQuery(start, end string) *DateRangeStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeStringInclusiveQuery(start, end string, startInclusive, endInclusive *bool) *DateRangeStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func (q *DateRangeStringQuery) SetBoost(b float64) { _ = "STUB: not implemented"; return }

func (q *DateRangeStringQuery) Boost() float64 { _ = "STUB: not implemented"; return 0 }

func (q *DateRangeStringQuery) SetField(f string) { _ = "STUB: not implemented"; return }

func (q *DateRangeStringQuery) Field() string { _ = "STUB: not implemented"; return "" }

func (q *DateRangeStringQuery) SetDateTimeParser(d string) { _ = "STUB: not implemented"; return }

func (q *DateRangeStringQuery) DateTimeParserName() string { _ = "STUB: not implemented"; return "" }

func (q *DateRangeStringQuery) Searcher(ctx context.Context, i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}

func (q *DateRangeStringQuery) parseEndpoints(startTime, endTime time.Time) (*float64, *float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (q *DateRangeStringQuery) Validate() error { _ = "STUB: not implemented"; return nil }

func isDateTimeWithinRange(t time.Time) bool { _ = "STUB: not implemented"; return false }
