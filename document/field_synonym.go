package document

import (
	"reflect"

	"github.com/blevesearch/bleve/v2/analysis"
	index "github.com/blevesearch/bleve_index_api"
)

var reflectStaticSizeSynonymField int

func init() {
	var f SynonymField
	reflectStaticSizeSynonymField = int(reflect.TypeOf(f).Size())
}

const DefaultSynonymIndexingOptions = index.IndexField

type SynonymField struct {
	name              string
	analyzer          analysis.Analyzer
	options           index.FieldIndexingOptions
	input             []string
	synonyms          []string
	numPlainTextBytes uint64

	synonymMap map[string][]string
}

func (s *SynonymField) Size() int { _ = "STUB: not implemented"; return 0 }

func (s *SynonymField) Name() string { _ = "STUB: not implemented"; return "" }

func (s *SynonymField) ArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (s *SynonymField) Options() index.FieldIndexingOptions {
	_ = "STUB: not implemented"
	return *new(index.FieldIndexingOptions)
}

func (s *SynonymField) NumPlainTextBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *SynonymField) AnalyzedLength() int { _ = "STUB: not implemented"; return 0 }

func (s *SynonymField) EncodedFieldType() byte { _ = "STUB: not implemented"; return 0 }

func (s *SynonymField) AnalyzedTokenFrequencies() index.TokenFrequencies {
	_ = "STUB: not implemented"
	return *new(index.TokenFrequencies)
}

func (s *SynonymField) Analyze() { _ = "STUB: not implemented"; return }

func (s *SynonymField) Value() []byte { _ = "STUB: not implemented"; return nil }

func (s *SynonymField) IterateSynonyms(visitor func(term string, synonyms []string)) {
	_ = "STUB: not implemented"
	return
}

func NewSynonymField(name string, analyzer analysis.Analyzer, input []string, synonyms []string) *SynonymField {
	_ = "STUB: not implemented"
	return nil
}

func processSynonymData(input []string, synonyms []string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func analyzeSynonymTerm(term string, analyzer analysis.Analyzer) string {
	_ = "STUB: not implemented"
	return ""
}
