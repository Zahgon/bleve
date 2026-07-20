package upsidedown

import (
	"sync"
)

type FieldCache struct {
	fieldIndexes   map[string]uint16
	indexFields    []string
	lastFieldIndex int
	mutex          sync.RWMutex
}

func NewFieldCache() *FieldCache { _ = "STUB: not implemented"; return nil }

func (f *FieldCache) AddExisting(field string, index uint16) { _ = "STUB: not implemented"; return }

func (f *FieldCache) addLOCKED(field string, index uint16) uint16 {
	_ = "STUB: not implemented"
	return 0
}

func (f *FieldCache) FieldNamed(field string, createIfMissing bool) (uint16, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (f *FieldCache) FieldIndexed(index uint16) (field string) {
	_ = "STUB: not implemented"
	return ""
}
