package scorer

import (
	"math"
)

var SqrtCache []float64

const MaxSqrtCache = 64

func init() {
	SqrtCache = make([]float64, MaxSqrtCache)
	for i := 0; i < MaxSqrtCache; i++ {
		SqrtCache[i] = math.Sqrt(float64(i))
	}
}
