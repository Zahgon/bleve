package document

import (
	"math"
	"reflect"
	"time"

	"github.com/blevesearch/bleve/v2/numeric"
	index "github.com/blevesearch/bleve_index_api"
)

var dateTimeValueSeperator = []byte{'\xff'}

var reflectStaticSizeDateTimeField int

func init() {
	var f DateTimeField
	reflectStaticSizeDateTimeField = int(reflect.TypeOf(f).Size())
}

const DefaultDateTimeIndexingOptions = index.StoreField | index.IndexField | index.DocValues
const DefaultDateTimePrecisionStep uint = 4

var MinTimeRepresentable = time.Unix(0, math.MinInt64)
var MaxTimeRepresentable = time.Unix(0, math.MaxInt64)

type DateTimeField struct {
	name              string
	arrayPositions    []uint64
	options           index.FieldIndexingOptions
	value             numeric.PrefixCoded
	numPlainTextBytes uint64
	length            int
	frequencies       index.TokenFrequencies
}

func (n *DateTimeField) Size() int { _ = "STUB: not implemented"; return 0 }

func (n *DateTimeField) Name() string { _ = "STUB: not implemented"; return "" }

func (n *DateTimeField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (n *DateTimeField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (n *DateTimeField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (n *DateTimeField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (n *DateTimeField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (n *DateTimeField) splitValue() (numeric.PrefixCoded, string) {
	_ = "STUB: not implemented"
	return *new(numeric.PrefixCoded), ""
}

func (n *DateTimeField) Analyze() { _ = "STUB: not implemented"; return }

func (n *DateTimeField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (n *DateTimeField) DateTime() (time.Time, string, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), "", nil
}

func (n *DateTimeField) GoString() string { _ = "STUB: not implemented"; return "" }

func (n *DateTimeField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func NewDateTimeFieldFromBytes(name string, arrayPositions []uint64, value []byte) *DateTimeField {
	_ = "STUB: not implemented"
	return nil
}

func NewDateTimeField(name string, arrayPositions []uint64, dt time.Time, layout string) (*DateTimeField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDateTimeFieldWithIndexingOptions(name string, arrayPositions []uint64, dt time.Time, layout string, options index.FieldIndexingOptions) (*DateTimeField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func canRepresent(dt time.Time) bool { _ = "STUB: not implemented"; return false }
