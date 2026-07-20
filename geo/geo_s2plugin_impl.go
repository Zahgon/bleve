package geo

import (
	"encoding/json"
	"sync"

	index "github.com/blevesearch/bleve_index_api"
	"github.com/blevesearch/geo/geojson"
	"github.com/blevesearch/geo/s2"
)

const (
	PointType              = "point"
	MultiPointType         = "multipoint"
	LineStringType         = "linestring"
	MultiLineStringType    = "multilinestring"
	PolygonType            = "polygon"
	MultiPolygonType       = "multipolygon"
	GeometryCollectionType = "geometrycollection"
	CircleType             = "circle"
	EnvelopeType           = "envelope"
)

var (
	spatialPluginsMap = make(map[string]index.SpatialAnalyzerPlugin)
	pluginsMapLock    = sync.RWMutex{}
)

func init() {
	registerS2RegionTermIndexer()
}

func registerS2RegionTermIndexer() { _ = "STUB: not implemented"; return }

func RegisterSpatialAnalyzerPlugin(plugin index.SpatialAnalyzerPlugin) {
	_ = "STUB: not implemented"
	return
}

func GetSpatialAnalyzerPlugin(typ string) index.SpatialAnalyzerPlugin {
	_ = "STUB: not implemented"
	return *new(index.SpatialAnalyzerPlugin)
}

func initS2IndexerOptions() s2.Options { _ = "STUB: not implemented"; return *new(s2.Options) }

func initS2SearcherOptions() s2.Options { _ = "STUB: not implemented"; return *new(s2.Options) }

func initS2OptionsForGeoPoints() s2.Options { _ = "STUB: not implemented"; return *new(s2.Options) }

type S2SpatialAnalyzerPlugin struct {
	s2Indexer                    *s2.RegionTermIndexer
	s2Searcher                   *s2.RegionTermIndexer
	s2GeoPointsRegionTermIndexer *s2.RegionTermIndexer
}

func (s *S2SpatialAnalyzerPlugin) Type() string { _ = "STUB: not implemented"; return "" }

func (s *S2SpatialAnalyzerPlugin) GetIndexTokens(queryShape index.GeoJSON) []string {
	_ = "STUB: not implemented"
	return nil
}

func (s *S2SpatialAnalyzerPlugin) GetQueryTokens(queryShape index.GeoJSON) []string {
	_ = "STUB: not implemented"
	return nil
}

type s2Tokenizable interface {
	IndexTokens(*s2.RegionTermIndexer) []string

	QueryTokens(*s2.RegionTermIndexer) []string
}

type s2TokenizableEx interface {
	IndexTokens(*S2SpatialAnalyzerPlugin) []string

	QueryTokens(*S2SpatialAnalyzerPlugin) []string
}

func (p *Point) Type() string { _ = "STUB: not implemented"; return "" }

func (p *Point) Value() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Point) Intersects(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Point) Contains(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Point) IndexTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

func (p *Point) QueryTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

type boundedRectangle struct {
	minLat float64
	maxLat float64
	minLon float64
	maxLon float64
}

func NewBoundedRectangle(minLat, minLon, maxLat,
	maxLon float64) *boundedRectangle {
	_ = "STUB: not implemented"
	return nil
}

func (br *boundedRectangle) Type() string { _ = "STUB: not implemented"; return "" }

func (br *boundedRectangle) Value() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *boundedRectangle) Intersects(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *boundedRectangle) Contains(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (br *boundedRectangle) IndexTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

func (br *boundedRectangle) QueryTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

type boundedPolygon struct {
	coordinates []Point
}

func NewBoundedPolygon(coordinates []Point) *boundedPolygon { _ = "STUB: not implemented"; return nil }

func (bp *boundedPolygon) Type() string { _ = "STUB: not implemented"; return "" }

func (bp *boundedPolygon) Value() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *boundedPolygon) Intersects(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *boundedPolygon) Contains(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (bp *boundedPolygon) IndexTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

func (bp *boundedPolygon) QueryTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

type pointDistance struct {
	dist      float64
	centerLat float64
	centerLon float64
}

func (p *pointDistance) Type() string { _ = "STUB: not implemented"; return "" }

func (p *pointDistance) Value() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewPointDistance(centerLat, centerLon,
	dist float64) *pointDistance {
	_ = "STUB: not implemented"
	return nil
}

func (p *pointDistance) Intersects(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pointDistance) Contains(s index.GeoJSON) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pd *pointDistance) IndexTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

func (pd *pointDistance) QueryTokens(s *S2SpatialAnalyzerPlugin) []string {
	_ = "STUB: not implemented"
	return nil
}

func NewGeometryCollection(coordinates [][][][][]float64,
	typs []string) (index.GeoJSON, []byte, error) {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON), nil, nil
}

func NewGeometryCollectionFromShapes(shapes []*geojson.GeoShape) (
	index.GeoJSON, []byte, error) {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON), nil, nil
}

func NewGeoCircleShape(cp []float64,
	radius string) (index.GeoJSON, []byte, error) {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON), nil, nil
}

func NewGeoJsonShape(coordinates [][][][]float64, typ string) (
	index.GeoJSON, []byte, error) {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON), nil, nil
}

func NewGeoJsonPoint(points []float64) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func NewGeoJsonMultiPoint(points [][]float64) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func NewGeoJsonLinestring(points [][]float64) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func NewGeoJsonMultilinestring(points [][][]float64) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func NewGeoJsonPolygon(points [][][]float64) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func NewGeoJsonMultiPolygon(points [][][][]float64) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func NewGeoCircle(points []float64, radius string) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func NewGeoEnvelope(points [][]float64) index.GeoJSON {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON)
}

func ParseGeoJSONShape(input json.RawMessage) (index.GeoJSON, error) {
	_ = "STUB: not implemented"
	return *new(index.GeoJSON), nil
}
