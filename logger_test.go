package log

import (
	"testing"

	"github.com/keyhorn/assert"
	"github.com/keyhorn/assert/matcher"
)

var equalTo = matcher.EqualTo
var endsWith = matcher.EndsWith

type TestWriter struct {
	Message string
}

func NewTestWriter() *TestWriter {
	return &TestWriter{
		Message: "",
	}
}

func (w *TestWriter) Open() error {
	return nil
}

func (w *TestWriter) Close() error {
	return nil
}

// Print calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Print.
func (w *TestWriter) Write(s string) error {
	w.Message = s
	return nil
}

func TestLogger(t *testing.T) {
	p := NewTestWriter()
	l := New(p)
	defer l.Close()
	
	l.Print("test1")
	assert.That(t, p.Message, endsWith("test1"))
	l.Println("test2")
	assert.That(t, p.Message, endsWith("test2\n"))
	l.Printf("test%v", 3)
	assert.That(t, p.Message, endsWith("test3"))
}
