package geo

import (
	"math"
)

var GeoBits uint = 32

var minLon = -180.0
var minLat = -90.0
var maxLon = 180.0
var maxLat = 90.0
var minLonRad = minLon * degreesToRadian
var minLatRad = minLat * degreesToRadian
var maxLonRad = maxLon * degreesToRadian
var maxLatRad = maxLat * degreesToRadian
var geoTolerance = 1e-6
var lonScale = float64((uint64(0x1)<<GeoBits)-1) / 360.0
var latScale = float64((uint64(0x1)<<GeoBits)-1) / 180.0

var geoHashMaxLength = 12

type Point struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

func MortonHash(lon, lat float64) uint64 { _ = "STUB: not implemented"; return 0 }

func scaleLon(lon float64) uint64 { _ = "STUB: not implemented"; return 0 }

func scaleLat(lat float64) uint64 { _ = "STUB: not implemented"; return 0 }

func MortonUnhashLon(hash uint64) float64 { _ = "STUB: not implemented"; return 0 }

func MortonUnhashLat(hash uint64) float64 { _ = "STUB: not implemented"; return 0 }

func unscaleLon(lon uint64) float64 { _ = "STUB: not implemented"; return 0 }

func unscaleLat(lat uint64) float64 { _ = "STUB: not implemented"; return 0 }

func compareGeo(a, b float64) float64 { _ = "STUB: not implemented"; return 0 }

func RectIntersects(aMinX, aMinY, aMaxX, aMaxY, bMinX, bMinY, bMaxX, bMaxY float64) bool {
	_ = "STUB: not implemented"
	return false
}

func RectWithin(aMinX, aMinY, aMaxX, aMaxY, bMinX, bMinY, bMaxX, bMaxY float64) bool {
	_ = "STUB: not implemented"
	return false
}

func BoundingBoxContains(lon, lat, minLon, minLat, maxLon, maxLat float64) bool {
	_ = "STUB: not implemented"
	return false
}

const degreesToRadian = math.Pi / 180
const radiansToDegrees = 180 / math.Pi

func DegreesToRadians(d float64) float64 { _ = "STUB: not implemented"; return 0 }

func RadiansToDegrees(r float64) float64 { _ = "STUB: not implemented"; return 0 }

var earthMeanRadiusMeters = 6371008.7714

func RectFromPointDistance(lon, lat, dist float64) (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func checkLatitude(latitude float64) error { _ = "STUB: not implemented"; return nil }

func checkLongitude(longitude float64) error { _ = "STUB: not implemented"; return nil }

func BoundingRectangleForPolygon(polygon []Point) (
	float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}
