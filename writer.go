package log

// Writer provides Writer interface.
type Writer interface {
	// Open opens the datasource for writing.
	Open() error
	// Close closes the datasource. 
	Close() error
	// Write writes text to the datasource.
	Write(string) error
}
