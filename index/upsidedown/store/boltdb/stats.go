package boltdb

type stats struct {
	s *Store
}

func (s *stats) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
