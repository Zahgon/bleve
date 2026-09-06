package util

var (
	BoltSnapshotsBucket           = []byte{'s'}
	BoltTrainerKey                = []byte{'t'}
	BoltTrainCompleteKey          = []byte{'c'}
	BoltTrainedSamplesKey         = []byte{'n'}
	BoltPathKey                   = []byte{'p'}
	BoltDeletedKey                = []byte{'d'}
	BoltInternalKey               = []byte{'i'}
	BoltMetaDataKey               = []byte{'m'}
	BoltMetaDataSegmentTypeKey    = []byte("type")
	BoltMetaDataSegmentVersionKey = []byte("version")
	BoltMetaDataTimeStamp         = []byte("timeStamp")
	BoltStatsKey                  = []byte("stats")
	BoltUpdatedFieldsKey          = []byte("fields")
	TotBytesWrittenKey            = []byte("TotBytesWritten")
	BoltMetaDataFileWriterIDKey   = []byte("fileWriterID")

	MappingInternalKey = []byte("_mapping")
)
