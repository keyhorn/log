package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// Logger is the interface for logging messages.
type Logger interface {
	// Printf writes a formated message to the log.
	Printf(format string, v ...interface{})
	// Print writes a message to the log.
	Print(v ...interface{})
	// Println writes a line to the log.
	Println(v ...interface{})
	// SetProvider changes provider.
	SetWriter(p Writer)
	// Close calls Provider.Close().
	Close()
}

// LoggerImpl represents an active logging object.
type LoggerImpl struct {
	mu     sync.Mutex        // ensures atomic writes; protects the folloing fields.
	writer Writer // log writer.
	closed bool
}

// New creates new Logger.
func New(w Writer) *LoggerImpl {
	l := &LoggerImpl{
		writer: w,
	}
	if err := l.writer.Open(); err != nil {
		return nil
	}
	return l
}

// Printf writes a formated message to the log.
func (l *LoggerImpl) Printf(format string, v ...interface{}) {
	l.output(2, fmt.Sprintf(format, v...))
}

// Print writes a message to the log.
func (l *LoggerImpl) Print(v ...interface{}) {
	l.output(2, fmt.Sprint(v...))
}

// Println writes a line to the log.
func (l *LoggerImpl) Println(v ...interface{}) {
	l.output(2, fmt.Sprintln(v...))
}

// SetWriter changes log writer implementation.
func (l *LoggerImpl) SetWriter(w Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.writer = w
}

// Close calls Provider.Close().
func (l *LoggerImpl) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.writer.Close()
	l.writer = nil
}

func (l *LoggerImpl) output(calldepth int, s string) error {
	now := time.Now()
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()

	// get file name and line no.
	l.mu.Unlock()
	var ok bool
	_, file, line, ok = runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	l.mu.Lock()

	text := fmt.Sprintf("%v %s:%d %s", now.Format("2006-01-02 15:04:05"), filepath.Base(file), line, s)

	return l.writer.Write(text)
}
