package scorch

import "time"

var RegistryAsyncErrorCallbacks = map[string]func(error, string){}

var RegistryEventCallbacks = map[string]func(Event) bool{}

type Event struct {
	Kind     EventKind
	Scorch   *Scorch
	Duration time.Duration
}

type EventKind int

const (
	EventKindCloseStart EventKind = iota

	EventKindClose

	EventKindMergerProgress

	EventKindPersisterProgress

	EventKindBatchIntroductionStart

	EventKindBatchIntroduction

	EventKindMergeTaskIntroductionStart

	EventKindMergeTaskIntroduction

	EventKindPreMergeCheck

	EventKindIndexStart

	EventKindPurgerCheck
)
