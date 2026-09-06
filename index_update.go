package bleve

import (
	"github.com/blevesearch/bleve/v2/mapping"
	index "github.com/blevesearch/bleve_index_api"
)

type pathInfo struct {
	fieldMapInfo []*fieldMapInfo
	dynamic      bool
	path         string
	analyser     string
	parentPath   string
}

type fieldMapInfo struct {
	fieldMapping   *mapping.FieldMapping
	analyzer       string
	datetimeParser string
	rootName       string
	parent         *pathInfo
}

func DeletedFields(ori, upd *mapping.IndexMappingImpl) (map[string]*index.UpdateFieldInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compareMappings(ori, upd *mapping.IndexMappingImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func checkUpdatedMapping(ori, upd *mapping.DocumentMapping) error {
	_ = "STUB: not implemented"
	return nil
}

func addPathInfo(paths map[string]*pathInfo, name string, mp *mapping.DocumentMapping,
	im *mapping.IndexMappingImpl, parent *pathInfo, rootName string) {
	_ = "STUB: not implemented"
	return
}

func compareCustomComponents(oriPaths, updPaths map[string]*pathInfo, ori, upd *mapping.IndexMappingImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func compareAnalysers(oriPaths, updPaths map[string]*pathInfo, ori, upd *mapping.IndexMappingImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func compareDateTimeParsers(oriPaths, updPaths map[string]*pathInfo, ori, upd *mapping.IndexMappingImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func compareSynonymSources(ori, upd *mapping.IndexMappingImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func compareAnalyserSubcomponents(ori, upd *mapping.IndexMappingImpl) error {
	_ = "STUB: not implemented"
	return nil
}

func addFieldInfo(fInfo map[string]*index.UpdateFieldInfo, ori, upd *pathInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func compareFieldMapping(original, updated *mapping.FieldMapping) (*index.UpdateFieldInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateFieldInfo(newInfo *index.UpdateFieldInfo, fInfo map[string]*index.UpdateFieldInfo,
	ori *pathInfo, oriFMapInfo *fieldMapInfo) error {
	_ = "STUB: not implemented"
	return nil
}
