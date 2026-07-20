package javadata

import (
	"bufio"
	"fmt"
	"io"
)

var ErrMalformedInput = fmt.Errorf("malformed input")

type Reader struct {
	r *bufio.Reader
}

func NewReader(r io.Reader) *Reader { _ = "STUB: not implemented"; return nil }

func (r *Reader) ReadBool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r *Reader) ReadInt32() (rv int32, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) ReadUint16() (rv uint16, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) ReadCharAsRune() (rv rune, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) ReadUTF() (string, error) { _ = "STUB: not implemented"; return "", nil }
