package utils

import (
	"bufio"
	"io"
)

type ReadWriter struct {
	*bufio.ReadWriter
}

func NewReadWriter(r io.Reader, w io.Writer) *ReadWriter {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	return &ReadWriter{
		ReadWriter: bufio.NewReadWriter(reader, writer),
	}
}

// func (rw *ReadWriter)
