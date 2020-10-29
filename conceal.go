package conceal

import (
	"io"
)

// A Writer processes the input bytes with the funcs it was instantiated with,
// then writes the result to the underlying io.Writer.
type Writer struct {
	w  io.Writer
	fs []func([]byte) []byte
}

// New creates a Writer with the provided underlying io.Writer and the processing funcs.
func New(w io.Writer, fs ...func([]byte) []byte) *Writer {
	return &Writer{w, fs}
}

func (w *Writer) Write(p []byte) (int, error) {
	return w.w.Write(w.apply(p))
}

func (w *Writer) apply(s []byte) []byte {
	for _, f := range w.fs {
		s = f(s)
	}
	return s
}
