package util

const DefaultFileCallbackId = ""

type FileWriter interface {
	Process(data []byte) []byte
	Id() string
}
type fileWriterImpl struct {
	id        string
	processor func(data []byte) []byte
}

func NewFileWriter(context []byte) (FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(FileWriter), nil
}

func (w *fileWriterImpl) Process(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func (w *fileWriterImpl) Id() string { _ = "STUB: not implemented"; return "" }

type FileReader interface {
	Process(data []byte) ([]byte, error)
	Id() string
}

type fileReaderImpl struct {
	id        string
	processor func(data []byte) ([]byte, error)
}

func NewFileReader(id string, context []byte) (FileReader, error) {
	_ = "STUB: not implemented"
	return *new(FileReader), nil
}

func (r *fileReaderImpl) Process(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *fileReaderImpl) Id() string { _ = "STUB: not implemented"; return "" }

var boltKeysProcessed = map[string]struct{}{
	string(BoltDeletedKey):       {},
	string(BoltInternalKey):      {},
	string(BoltStatsKey):         {},
	string(BoltUpdatedFieldsKey): {},
}
