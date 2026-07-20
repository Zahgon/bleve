package bleve

import (
	"sync"
)

type IndexStat struct {
	searches   uint64
	searchTime uint64
	i          *indexImpl
}

func (is *IndexStat) statsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (is *IndexStat) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type IndexStats struct {
	indexes map[string]*IndexStat
	mutex   sync.RWMutex
}

func NewIndexStats() *IndexStats { _ = "STUB: not implemented"; return nil }

func (i *IndexStats) Register(index Index) { _ = "STUB: not implemented"; return }

func (i *IndexStats) UnRegister(index Index) { _ = "STUB: not implemented"; return }

func (i *IndexStats) String() string { _ = "STUB: not implemented"; return "" }

var indexStats *IndexStats
