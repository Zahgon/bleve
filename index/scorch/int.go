package scorch

const (
	intMin      = 0x80
	intMaxWidth = 8
	intZero     = intMin + intMaxWidth
	intSmall    = intMax - intZero - intMaxWidth

	intMax = 0xfd
)

func encodeUvarintAscending(b []byte, v uint64) []byte { _ = "STUB: not implemented"; return nil }

func decodeUvarintAscending(b []byte) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
