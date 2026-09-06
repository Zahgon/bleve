package geo

import (
	"github.com/blevesearch/geo/geojson"
)

func ExtractGeoPoint(thing interface{}) (lon, lat float64, success bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

type loner interface {
	Lon() float64
}

type later interface {
	Lat() float64
}

type lnger interface {
	Lng() float64
}

var GlueBytes = []byte("##")

var GlueBytesOffset = len(GlueBytes)

func extractCoordinates(thing interface{}) []float64 { _ = "STUB: not implemented"; return nil }

func extract2DCoordinates(thing interface{}) [][]float64 { _ = "STUB: not implemented"; return nil }

func extract3DCoordinates(thing interface{}) (c [][][]float64) {
	_ = "STUB: not implemented"
	return nil
}

func extract4DCoordinates(thing interface{}) (rv [][][][]float64) {
	_ = "STUB: not implemented"
	return nil
}

func ParseGeoShapeField(thing interface{}) (interface{}, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func extractGeoShape(thing interface{}) (*geojson.GeoShape, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func ExtractGeometryCollection(thing interface{}) ([]*geojson.GeoShape, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func ExtractCircle(thing interface{}) (*geojson.GeoShape, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func ExtractGeoShapeCoordinates(coordValue interface{},
	typ string) (*geojson.GeoShape, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
