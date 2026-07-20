package camelcase

type State interface {
	StartSym(sym rune) bool

	Member(sym rune, peek *rune) bool
}

type LowerCaseState struct{}

func (s *LowerCaseState) Member(sym rune, peek *rune) bool { _ = "STUB: not implemented"; return false }

func (s *LowerCaseState) StartSym(sym rune) bool { _ = "STUB: not implemented"; return false }

type UpperCaseState struct {
	startedCollecting bool
	collectingUpper   bool
}

func (s *UpperCaseState) Member(sym rune, peek *rune) bool { _ = "STUB: not implemented"; return false }

func (s *UpperCaseState) StartSym(sym rune) bool { _ = "STUB: not implemented"; return false }

type NumberCaseState struct{}

func (s *NumberCaseState) Member(sym rune, peek *rune) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *NumberCaseState) StartSym(sym rune) bool { _ = "STUB: not implemented"; return false }

type NonAlphaNumericCaseState struct{}

func (s *NonAlphaNumericCaseState) Member(sym rune, peek *rune) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *NonAlphaNumericCaseState) StartSym(sym rune) bool { _ = "STUB: not implemented"; return false }
