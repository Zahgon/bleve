package bleve

import (
	"time"

	"github.com/blevesearch/bleve/v2/search/query"
)

func NewBoolFieldQuery(val bool) *query.BoolFieldQuery { _ = "STUB: not implemented"; return nil }

func NewBooleanQuery() *query.BooleanQuery { _ = "STUB: not implemented"; return nil }

func NewConjunctionQuery(conjuncts ...query.Query) *query.ConjunctionQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeQuery(start, end time.Time) *query.DateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeInclusiveQuery(start, end time.Time, startInclusive, endInclusive *bool) *query.DateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeStringQuery(start, end string) *query.DateRangeStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeInclusiveStringQuery(start, end string, startInclusive, endInclusive *bool) *query.DateRangeStringQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewDisjunctionQuery(disjuncts ...query.Query) *query.DisjunctionQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewDocIDQuery(ids []string) *query.DocIDQuery { _ = "STUB: not implemented"; return nil }

func NewFuzzyQuery(term string) *query.FuzzyQuery { _ = "STUB: not implemented"; return nil }

func NewMatchAllQuery() *query.MatchAllQuery { _ = "STUB: not implemented"; return nil }

func NewMatchNoneQuery() *query.MatchNoneQuery { _ = "STUB: not implemented"; return nil }

func NewMatchPhraseQuery(matchPhrase string) *query.MatchPhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewMatchQuery(match string) *query.MatchQuery { _ = "STUB: not implemented"; return nil }

func NewNumericRangeQuery(min, max *float64) *query.NumericRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewNumericRangeInclusiveQuery(min, max *float64, minInclusive, maxInclusive *bool) *query.NumericRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewTermRangeQuery(min, max string) *query.TermRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewTermRangeInclusiveQuery(min, max string, minInclusive, maxInclusive *bool) *query.TermRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewPhraseQuery(terms []string, field string) *query.PhraseQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewPrefixQuery(prefix string) *query.PrefixQuery { _ = "STUB: not implemented"; return nil }

func NewRegexpQuery(regexp string) *query.RegexpQuery { _ = "STUB: not implemented"; return nil }

func NewQueryStringQuery(q string) *query.QueryStringQuery { _ = "STUB: not implemented"; return nil }

func NewTermQuery(term string) *query.TermQuery { _ = "STUB: not implemented"; return nil }

func NewWildcardQuery(wildcard string) *query.WildcardQuery { _ = "STUB: not implemented"; return nil }

func NewGeoBoundingBoxQuery(topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64) *query.GeoBoundingBoxQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoDistanceQuery(lon, lat float64, distance string) *query.GeoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}

func NewIPRangeQuery(cidr string) *query.IPRangeQuery { _ = "STUB: not implemented"; return nil }

func NewGeoShapeQuery(coordinates [][][][]float64, typ, relation string) (*query.GeoShapeQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeoShapeCircleQuery(coordinates []float64, radius, relation string) (*query.GeoShapeQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeometryCollectionQuery(coordinates [][][][][]float64, types []string, relation string) (*query.GeoShapeQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
