package registry

import (
	"sync"

	"github.com/blevesearch/bleve/v2/search"
)

type NestedFieldCache struct {
	prefixDepth map[string]int
	once        sync.Once
	m           sync.RWMutex
}

func NewNestedFieldCache() *NestedFieldCache { _ = "STUB: not implemented"; return nil }

func (nfc *NestedFieldCache) InitOnce(buildFunc func() map[string]int) {
	_ = "STUB: not implemented"
	return
}

func (nfc *NestedFieldCache) NestedDepth(fieldPaths search.FieldSet) (common int, max int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (nfc *NestedFieldCache) CountNested() int { _ = "STUB: not implemented"; return 0 }

func (nfc *NestedFieldCache) IntersectsPrefix(fieldPaths search.FieldSet) bool {
	_ = "STUB: not implemented"
	return false
}

func (nfc *NestedFieldCache) prefixMatch(prefix string, fieldPaths search.FieldSet) (common bool, any bool) {
	_ = "STUB: not implemented"
	return false, false
}
