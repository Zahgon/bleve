package registry

import (
	"github.com/blevesearch/bleve/v2/analysis"
)

func RegisterSynonymSource(typ string, constructor SynonymSourceConstructor) error {
	_ = "STUB: not implemented"
	return nil
}

type SynonymSourceCache struct {
	*ConcurrentCache
}

func NewSynonymSourceCache() *SynonymSourceCache { _ = "STUB: not implemented"; return nil }

type SynonymSourceConstructor func(config map[string]interface{}, cache *Cache) (analysis.SynonymSource, error)
type SynonymSourceRegistry map[string]SynonymSourceConstructor

func SynonymSourceBuild(name string, config map[string]interface{}, cache *Cache) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *SynonymSourceCache) SynonymSourceNamed(name string, cache *Cache) (analysis.SynonymSource, error) {
	_ = "STUB: not implemented"
	return *new(analysis.SynonymSource), nil
}

func (c *SynonymSourceCache) DefineSynonymSource(name string, typ string, config map[string]interface{}, cache *Cache) (analysis.SynonymSource, error) {
	_ = "STUB: not implemented"
	return *new(analysis.SynonymSource), nil
}

func (c *SynonymSourceCache) VisitSynonymSources(visitor analysis.SynonymSourceVisitor) error {
	_ = "STUB: not implemented"
	return nil
}
