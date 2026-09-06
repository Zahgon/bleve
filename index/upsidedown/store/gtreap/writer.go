package gtreap

type Writer struct {
	s *Store
}

func (w *Writer) NewBatch() store.KVBatch { _ = "STUB: not implemented"; return *new(store.KVBatch) }

func (w *Writer) NewBatchEx(options store.KVBatchOptions) ([]byte, store.KVBatch, error) {
	_ = "STUB: not implemented"
	return nil, *new(store.KVBatch), nil
}

func (w *Writer) ExecuteBatch(batch store.KVBatch) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }
