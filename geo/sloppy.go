package geo

import (
	"math"
)

var earthDiameterPerLatitude []float64

const (
	radiusTabsSize = (1 << 10) + 1
	radiusDelta    = (math.Pi / 2) / (radiusTabsSize - 1)
	radiusIndexer  = 1 / radiusDelta
)

func init() {

	a := 6378137.0
	b := 6356752.31420
	a2 := a * a
	b2 := b * b
	earthDiameterPerLatitude = make([]float64, radiusTabsSize)
	earthDiameterPerLatitude[0] = 2.0 * a / 1000
	earthDiameterPerLatitude[radiusTabsSize-1] = 2.0 * b / 1000
	for i := 1; i < radiusTabsSize-1; i++ {
		lat := math.Pi * float64(i) / (2*radiusTabsSize - 1)
		one := math.Pow(a2*math.Cos(lat), 2)
		two := math.Pow(b2*math.Sin(lat), 2)
		three := math.Pow(float64(a)*math.Cos(lat), 2)
		four := math.Pow(b*math.Sin(lat), 2)
		radius := math.Sqrt((one + two) / (three + four))
		earthDiameterPerLatitude[i] = 2 * radius / 1000
	}
}

func earthDiameter(lat float64) float64 { _ = "STUB: not implemented"; return 0 }
