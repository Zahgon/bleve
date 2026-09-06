package upsidedown

import (
	"encoding/binary"
)

var mergeOperator upsideDownMerge

var dictionaryTermIncr []byte
var dictionaryTermDecr []byte

func init() {
	dictionaryTermIncr = make([]byte, 8)
	binary.LittleEndian.PutUint64(dictionaryTermIncr, uint64(1))
	dictionaryTermDecr = make([]byte, 8)
	var negOne = int64(-1)
	binary.LittleEndian.PutUint64(dictionaryTermDecr, uint64(negOne))
}

type upsideDownMerge struct{}

func (m *upsideDownMerge) FullMerge(key, existingValue []byte, operands [][]byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *upsideDownMerge) PartialMerge(key, leftOperand, rightOperand []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *upsideDownMerge) Name() string { _ = "STUB: not implemented"; return "" }
