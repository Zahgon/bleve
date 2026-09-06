package scorch

import (
	index "github.com/blevesearch/bleve_index_api"
)

var OptimizeConjunction = true
var OptimizeConjunctionUnadorned = true
var OptimizeDisjunctionUnadorned = true

func (s *IndexSnapshotTermFieldReader) Optimize(kind string,
	octx index.OptimizableContext) (index.OptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.OptimizableContext), nil
}

var OptimizeDisjunctionUnadornedMinChildCardinality = uint64(256)

func (s *IndexSnapshotTermFieldReader) optimizeConjunction(
	octx index.OptimizableContext) (index.OptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.OptimizableContext), nil
}

type OptimizeTFRConjunction struct {
	snapshot *IndexSnapshot

	tfrs []*IndexSnapshotTermFieldReader
}

func (o *OptimizeTFRConjunction) Finish() (index.Optimized, error) {
	_ = "STUB: not implemented"
	return *new(index.Optimized), nil
}

func (s *IndexSnapshotTermFieldReader) optimizeConjunctionUnadorned(
	octx index.OptimizableContext) (index.OptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.OptimizableContext), nil
}

type OptimizeTFRConjunctionUnadorned struct {
	snapshot *IndexSnapshot

	tfrs []*IndexSnapshotTermFieldReader
}

var OptimizeTFRConjunctionUnadornedTerm = []byte("<conjunction:unadorned>")
var OptimizeTFRConjunctionUnadornedField = "*"

func (o *OptimizeTFRConjunctionUnadorned) Finish() (rv index.Optimized, err error) {
	_ = "STUB: not implemented"
	return *new(index.Optimized), nil
}

func (s *IndexSnapshotTermFieldReader) optimizeDisjunctionUnadorned(
	octx index.OptimizableContext) (index.OptimizableContext, error) {
	_ = "STUB: not implemented"
	return *new(index.OptimizableContext), nil
}

type OptimizeTFRDisjunctionUnadorned struct {
	snapshot *IndexSnapshot

	tfrs []*IndexSnapshotTermFieldReader
}

var OptimizeTFRDisjunctionUnadornedTerm = []byte("<disjunction:unadorned>")
var OptimizeTFRDisjunctionUnadornedField = "*"

func (o *OptimizeTFRDisjunctionUnadorned) Finish() (rv index.Optimized, err error) {
	_ = "STUB: not implemented"
	return *new(index.Optimized), nil
}

func (i *IndexSnapshot) unadornedTermFieldReader(
	term []byte, field string) *IndexSnapshotTermFieldReader {
	_ = "STUB: not implemented"
	return nil
}
