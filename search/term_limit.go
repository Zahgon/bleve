package search

import (
	"context"
	"sync/atomic"
)

type termSearchersCounter struct {
	limit int
	count atomic.Uint64
}

func ContextWithTermSearchersCounter(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func RecordTermSearcher(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
