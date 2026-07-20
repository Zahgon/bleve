package upsidedown

type indexStat struct {
	updates, deletes, batches, errors uint64
	analysisTime, indexTime           uint64
	termSearchersStarted              uint64
	termSearchersFinished             uint64
	numPlainTextBytesIndexed          uint64
	i                                 *UpsideDownCouch
}

func (i *indexStat) statsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (i *indexStat) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
