package bleve

const (
	ErrorIndexPathExists Error = iota
	ErrorIndexPathDoesNotExist
	ErrorIndexMetaMissing
	ErrorIndexMetaCorrupt
	ErrorIndexClosed
	ErrorAliasMulti
	ErrorAliasEmpty
	ErrorUnknownIndexType
	ErrorEmptyID
	ErrorIndexReadInconsistency
	ErrorTwoPhaseSearchInconsistency
	ErrorSynonymSearchNotSupported
	ErrorTrainingNotSupported
)

type Error int

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

var errorMessages = map[Error]string{
	ErrorIndexPathExists:             "cannot create new index, path already exists",
	ErrorIndexPathDoesNotExist:       "cannot open index, path does not exist",
	ErrorIndexMetaMissing:            "cannot open index, metadata missing",
	ErrorIndexMetaCorrupt:            "cannot open index, metadata corrupt",
	ErrorIndexClosed:                 "index is closed",
	ErrorAliasMulti:                  "cannot perform single index operation on multiple index alias",
	ErrorAliasEmpty:                  "cannot perform operation on empty alias",
	ErrorUnknownIndexType:            "unknown index type",
	ErrorEmptyID:                     "document ID cannot be empty",
	ErrorIndexReadInconsistency:      "index read inconsistency detected",
	ErrorTwoPhaseSearchInconsistency: "2-phase search failed, likely due to an overlapping topology change",
	ErrorSynonymSearchNotSupported:   "synonym search not supported",
}
