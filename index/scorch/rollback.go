package scorch

type RollbackPoint struct {
	epoch uint64
	meta  map[string][]byte
}

func (r *RollbackPoint) GetInternal(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func RollbackPoints(path string) ([]*RollbackPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Rollback(path string, to *RollbackPoint) error { _ = "STUB: not implemented"; return nil }
