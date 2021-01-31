package log

import (
	"fmt"
	"os"
)

// StdoutWriter is stdout writer.
type StdoutWriter struct {
	stdout *os.File
}

// NewStdoutWriter creates stdout writer.
func NewStdoutWriter() *StdoutWriter {
	return &StdoutWriter{}
}

// Open opens the datasource for writing.
func (w *StdoutWriter) Open() error {
	w.stdout = os.Stdout
	return nil
}

// Close closes the datasource.
func Close() error {
	// nothing to do.
	return nil
}

// Write writes text to the datasource.
func (w *StdoutWriter) Write(s string) error {
	_, err := fmt.Fprint(w.stdout, s)
	return err
}
