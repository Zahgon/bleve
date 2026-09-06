package moss

type stats struct {
	s *Store
}

func (s *stats) statsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (s *stats) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
