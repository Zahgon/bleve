package mapping

import (
	"io"
	"log"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/document"
	"github.com/blevesearch/bleve/v2/search"
)

type Classifier interface {
	Type() string
}

type bleveClassifier interface {
	BleveType() string
}

var logger = log.New(io.Discard, "bleve mapping ", log.LstdFlags)

func SetLog(l *log.Logger) { _ = "STUB: not implemented"; return }

type IndexMapping interface {
	MapDocument(doc *document.Document, data interface{}) error
	Validate() error

	DateTimeParserNamed(name string) analysis.DateTimeParser

	DefaultSearchField() string

	AnalyzerNameForPath(path string) string
	AnalyzerNamed(name string) analysis.Analyzer

	FieldMappingForPath(path string) FieldMapping
}

type SynonymMapping interface {
	IndexMapping

	MapSynonymDocument(doc *document.Document, collection string, input []string, synonyms []string) error

	SynonymSourceForPath(path string) string

	SynonymSourceNamed(name string) analysis.SynonymSource

	SynonymCount() int

	SynonymSourceVisitor(visitor analysis.SynonymSourceVisitor) error
}

type NestedMapping interface {
	NestedDepth(fieldPaths search.FieldSet) (int, int)

	IntersectsPrefix(fieldPaths search.FieldSet) bool

	CountNested() int
}
