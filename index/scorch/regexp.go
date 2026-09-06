package scorch

import (
	"regexp/syntax"

	"github.com/blevesearch/vellum/regexp"
)

func parseRegexp(pattern string) (a *regexp.Regexp, prefixBeg, prefixEnd []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func literalPrefix(s *syntax.Regexp) string { _ = "STUB: not implemented"; return "" }
