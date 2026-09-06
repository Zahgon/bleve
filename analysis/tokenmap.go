package analysis

type TokenMap map[string]bool

func NewTokenMap() TokenMap { _ = "STUB: not implemented"; return *new(TokenMap) }

func (t TokenMap) LoadFile(filename string) error { _ = "STUB: not implemented"; return nil }

func (t TokenMap) LoadBytes(data []byte) error { _ = "STUB: not implemented"; return nil }

func (t TokenMap) LoadLine(line string) { _ = "STUB: not implemented"; return }

func (t TokenMap) AddToken(token string) { _ = "STUB: not implemented"; return }
