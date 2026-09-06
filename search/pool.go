package search

import (
	"reflect"
)

var reflectStaticSizeDocumentMatchPool int

func init() {
	var dmp DocumentMatchPool
	reflectStaticSizeDocumentMatchPool = int(reflect.TypeOf(dmp).Size())
}

type DocumentMatchPoolTooSmall func(p *DocumentMatchPool) *DocumentMatch

type DocumentMatchPool struct {
	avail    DocumentMatchCollection
	TooSmall DocumentMatchPoolTooSmall
}

func defaultDocumentMatchPoolTooSmall(p *DocumentMatchPool) *DocumentMatch {
	_ = "STUB: not implemented"
	return nil
}

func NewDocumentMatchPool(size, sortsize int) *DocumentMatchPool {
	_ = "STUB: not implemented"
	return nil
}

func (p *DocumentMatchPool) Get() *DocumentMatch { _ = "STUB: not implemented"; return nil }

func (p *DocumentMatchPool) Put(d *DocumentMatch) { _ = "STUB: not implemented"; return }
