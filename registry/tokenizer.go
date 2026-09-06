package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func RegisterTokenizer(name string, constructor TokenizerConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type TokenizerConstructor func(config map[string]interface{}, cache *Cache) (analysis.Tokenizer, error)
type TokenizerRegistry map[string]TokenizerConstructor

type TokenizerCache struct {
	*ConcurrentCache
}

func NewTokenizerCache() *TokenizerCache { _ = "STUB: not implemented"; return nil }

func TokenizerBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TokenizerCache) TokenizerNamed(name string, cache *Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func (c *TokenizerCache) DefineTokenizer(name string, typ string, config map[string]interface{}, cache *Cache) (analysis.Tokenizer, error) {
	_ = "STUB: not implemented"
	return *new(analysis.Tokenizer), nil
}

func TokenizerTypesAndInstances() ([]string, []string) { _ = "STUB: not implemented"; return nil, nil }
