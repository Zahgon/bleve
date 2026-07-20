package numeric

const ShiftStartInt64 byte = 0x20

type PrefixCoded []byte

func NewPrefixCodedInt64(in int64, shift uint) (PrefixCoded, error) {
	_ = "STUB: not implemented"
	return *new(PrefixCoded), nil
}

func NewPrefixCodedInt64Prealloc(in int64, shift uint, prealloc []byte) (
	rv PrefixCoded, preallocRest []byte, err error) {
	_ = "STUB: not implemented"
	return *new(PrefixCoded), nil, nil
}

func MustNewPrefixCodedInt64(in int64, shift uint) PrefixCoded {
	_ = "STUB: not implemented"
	return *new(PrefixCoded)
}

func MustNewPrefixCodedInt64Prealloc(in int64, shift uint, prealloc []byte) PrefixCoded {
	_ = "STUB: not implemented"
	return *new(PrefixCoded)
}

func (p PrefixCoded) Shift() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (p PrefixCoded) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func ValidPrefixCodedTerm(p string) (bool, int) { _ = "STUB: not implemented"; return false, 0 }

func ValidPrefixCodedTermBytes(p []byte) (bool, int) { _ = "STUB: not implemented"; return false, 0 }
