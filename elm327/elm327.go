package elm327

import (
	"bufio"
	"errors"
	"io"
)

var ErrUnknownCommand = errors.New("invalid or unknown command")

type stringReader interface {
	ReadString(delim byte) (string, error)
}
type ELM327 struct {
	r    stringReader
	w    io.Writer
	last string
}

func New(rw io.ReadWriter) *ELM327 {
	c := &ELM327{w: rw}
	if r, ok := rw.(stringReader); ok {
		c.r = r
	} else {
		c.r = bufio.NewReader(rw)
	}
	return c
}
