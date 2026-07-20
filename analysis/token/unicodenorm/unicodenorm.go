package unicodenorm

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"golang.org/x/text/unicode/norm"
)

const Name = "normalize_unicode"

const NFC = "nfc"
const NFD = "nfd"
const NFKC = "nfkc"
const NFKD = "nfkd"

var forms = map[string]norm.Form{
	NFC:  norm.NFC,
	NFD:  norm.NFD,
	NFKC: norm.NFKC,
	NFKD: norm.NFKD,
}

type UnicodeNormalizeFilter struct {
	form norm.Form
}

func NewUnicodeNormalizeFilter(formName string) (*UnicodeNormalizeFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MustNewUnicodeNormalizeFilter(formName string) *UnicodeNormalizeFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *UnicodeNormalizeFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	_ = "STUB: not implemented"
	return *new(analysis.TokenStream)
}

func UnicodeNormalizeFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func init() {
	err := registry.RegisterTokenFilter(Name, UnicodeNormalizeFilterConstructor)
	if err != nil {
		panic(err)
	}
}
