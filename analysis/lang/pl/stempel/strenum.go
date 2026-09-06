package stempel

type strEnum struct {
	r    []rune
	from int
	by   int
}

func newStrEnum(s []rune, up bool) *strEnum { _ = "STUB: not implemented"; return nil }

func (s *strEnum) next() (rune, error) { _ = "STUB: not implemented"; return 0, nil }
