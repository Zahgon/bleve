package search

import (
	"reflect"
)

var reflectStaticSizeExplanation int

func init() {
	var e Explanation
	reflectStaticSizeExplanation = int(reflect.TypeOf(e).Size())
}

const MergedExplMessage = "sum of merged explanations:"

type Explanation struct {
	Value        float64        `json:"value"`
	Message      string         `json:"message"`
	PartialMatch bool           `json:"partial_match,omitempty"`
	Children     []*Explanation `json:"children,omitempty"`
}

func (expl *Explanation) String() string { _ = "STUB: not implemented"; return "" }

func (expl *Explanation) Size() int { _ = "STUB: not implemented"; return 0 }

func (expl *Explanation) MergeWith(other *Explanation) *Explanation {
	_ = "STUB: not implemented"
	return nil
}
