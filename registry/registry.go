package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/search/highlight"
)

var stores = make(KVStoreRegistry, 0)
var indexTypes = make(IndexTypeRegistry, 0)

var fragmentFormatters = make(FragmentFormatterRegistry, 0)
var fragmenters = make(FragmenterRegistry, 0)
var highlighters = make(HighlighterRegistry, 0)

var charFilters = make(CharFilterRegistry, 0)
var tokenizers = make(TokenizerRegistry, 0)
var tokenMaps = make(TokenMapRegistry, 0)
var tokenFilters = make(TokenFilterRegistry, 0)
var analyzers = make(AnalyzerRegistry, 0)
var dateTimeParsers = make(DateTimeParserRegistry, 0)
var synonymSources = make(SynonymSourceRegistry, 0)

type Cache struct {
	CharFilters        *CharFilterCache
	Tokenizers         *TokenizerCache
	TokenMaps          *TokenMapCache
	TokenFilters       *TokenFilterCache
	Analyzers          *AnalyzerCache
	DateTimeParsers    *DateTimeParserCache
	FragmentFormatters *FragmentFormatterCache
	Fragmenters        *FragmenterCache
	Highlighters       *HighlighterCache
	SynonymSources     *SynonymSourceCache
	NestedPrefixes     *NestedFieldCache
}

func NewCache() *Cache { _ = "STUB: not implemented"; return nil }

func typeFromConfig(config map[string]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Cache) CharFilterNamed(name string) (analysis.CharFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.CharFilter), nil
}

func (c *Cache) DefineCharFilter(name string, config map[string]interface{}) (analysis.CharFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.CharFilter), nil
}

func (c *Cache) TokenizerNamed(name string) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func (c *Cache) DefineTokenizer(name string, config map[string]interface{}) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func (c *Cache) TokenMapNamed(name string) (analysis.TokenMap, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenMap), nil
}

func (c *Cache) DefineTokenMap(name string, config map[string]interface{}) (analysis.TokenMap, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenMap), nil
}

func (c *Cache) TokenFilterNamed(name string) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func (c *Cache) DefineTokenFilter(name string, config map[string]interface{}) (analysis.TokenFilter, error) {
	_ = "STUB: not implemented"
	return *new(analysis.TokenFilter), nil
}

func (c *Cache) AnalyzerNamed(name string) (analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer), nil
}

func (c *Cache) DefineAnalyzer(name string, config map[string]interface{}) (analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Analyzer), nil
}

func (c *Cache) DateTimeParserNamed(name string) (analysis.DateTimeParser, error) {
	_ = "STUB: not implemented"
	return *new(analysis.DateTimeParser), nil
}

func (c *Cache) DefineDateTimeParser(name string, config map[string]interface{}) (analysis.DateTimeParser, error) {
	_ = "STUB: not implemented"
	return *new(analysis.DateTimeParser), nil
}

func (c *Cache) SynonymSourceNamed(name string) (analysis.SynonymSource, error) {
	_ = "STUB: not implemented"
	return *new(analysis.SynonymSource), nil
}

func (c *Cache) DefineSynonymSource(name string, config map[string]interface{}) (analysis.SynonymSource, error) {
	_ = "STUB: not implemented"
	return *new(analysis.SynonymSource), nil
}

func (c *Cache) FragmentFormatterNamed(name string) (highlight.FragmentFormatter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.FragmentFormatter), nil
}

func (c *Cache) DefineFragmentFormatter(name string, config map[string]interface{}) (highlight.FragmentFormatter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.FragmentFormatter), nil
}

func (c *Cache) FragmenterNamed(name string) (highlight.Fragmenter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Fragmenter), nil
}

func (c *Cache) DefineFragmenter(name string, config map[string]interface{}) (highlight.Fragmenter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Fragmenter), nil
}

func (c *Cache) HighlighterNamed(name string) (highlight.Highlighter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Highlighter), nil
}

func (c *Cache) DefineHighlighter(name string, config map[string]interface{}) (highlight.Highlighter, error) {
	_ = "STUB: not implemented"
	return *new(highlight.Highlighter), nil
}
