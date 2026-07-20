package mergeplan

import (
	"errors"
)

type Segment interface {
	Id() uint64

	FullSize() int64

	LiveSize() int64

	HasVector() bool

	FileSize() int64

	LiveFileSize() int64
}

func Plan(segments []Segment, o *MergePlanOptions) (*MergePlan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MergePlan struct {
	Tasks []*MergeTask
}

type MergeTask struct {
	Segments []Segment
}

type MergePlanOptions struct {
	MaxSegmentsPerTier int

	MaxSegmentSize int64

	MaxSegmentFileSize int64

	TierGrowth float64

	SegmentsPerMergeTask int

	FloorSegmentSize int64

	FloorSegmentFileSize int64

	ReclaimDeletesWeight float64

	CalcBudget func(totalSize int64, firstTierSize int64,
		o *MergePlanOptions) (budgetNumSegments int)

	ScoreSegments func(segments []Segment, o *MergePlanOptions) float64

	Logger func(string)
}

func (o *MergePlanOptions) RaiseToFloorSegmentSize(s int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (o *MergePlanOptions) RaiseToFloorSegmentFileSize(s int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (o *MergePlanOptions) BudgetCurrency() BudgetCurrency {
	_ = "STUB: not implemented"
	return *new(BudgetCurrency)
}

const MaxSegmentSizeLimit = 1<<31 - 1

var ErrMaxSegmentSizeTooLarge = errors.New("MaxSegmentSize exceeds the size limit")

var DefaultMergePlanOptions = MergePlanOptions{
	MaxSegmentsPerTier:   10,
	MaxSegmentSize:       5000000,
	MaxSegmentFileSize:   4 * 1024 * 1024 * 1024,
	TierGrowth:           10.0,
	SegmentsPerMergeTask: 10,
	FloorSegmentSize:     2000,
	ReclaimDeletesWeight: 2.0,
}

var SingleSegmentMergePlanOptions = MergePlanOptions{
	MaxSegmentsPerTier:   1,
	MaxSegmentSize:       1 << 30,
	MaxSegmentFileSize:   1 << 40,
	TierGrowth:           1.0,
	SegmentsPerMergeTask: 10,
	FloorSegmentSize:     1 << 30,
	ReclaimDeletesWeight: 2.0,
	FloorSegmentFileSize: 1 << 40,
}

type BudgetCurrency int

const (
	FileSizeBudget BudgetCurrency = iota

	LiveSizeBudget
)

var ErrUnknownCurrency = errors.New("unknown budget currency")

type rosterCandidate struct {
	segments []Segment
	score    float64
}

func (c *rosterCandidate) betterThan(other *rosterCandidate) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *rosterCandidate) count() int { _ = "STUB: not implemented"; return 0 }

func plan(segmentsIn []Segment, o *MergePlanOptions) (*MergePlan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CalcBudget(totalSize int64, firstTierSize int64, o *MergePlanOptions) (
	budgetNumSegments int) {
	_ = "STUB: not implemented"
	return 0
}

func removeSegments(segments []Segment, toRemove []Segment) []Segment {
	_ = "STUB: not implemented"
	return nil
}

func ScoreSegments(segments []Segment, o *MergePlanOptions) float64 {
	_ = "STUB: not implemented"
	return 0
}

func ToBarChart(prefix string, barMax int, segments []Segment, plan *MergePlan) string {
	_ = "STUB: not implemented"
	return ""
}

func ValidateMergePlannerOptions(options *MergePlanOptions) error {
	_ = "STUB: not implemented"
	return nil
}
