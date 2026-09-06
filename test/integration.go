package test

import (
	"github.com/blevesearch/bleve/v2"
)

type SearchTest struct {
	Search  *bleve.SearchRequest `json:"search"`
	Result  *bleve.SearchResult  `json:"result"`
	Comment string               `json:"comment"`
}

type SearchTests []*SearchTest
