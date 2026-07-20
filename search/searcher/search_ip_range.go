package searcher

import (
	"context"
	"net"

	"github.com/blevesearch/bleve/v2/search"
	index "github.com/blevesearch/bleve_index_api"
)

func netLimits(n *net.IPNet) (lo net.IP, hi net.IP) {
	_ = "STUB: not implemented"
	return *new(net.IP), *new(net.IP)
}

func NewIPRangeSearcher(ctx context.Context, indexReader index.IndexReader, ipNet *net.IPNet,
	field string, boost float64, options search.SearcherOptions) (
	search.Searcher, error) {
	_ = "STUB: not implemented"
	return *new(search.Searcher), nil
}
