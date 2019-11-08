package ulog

import (
	"fmt"
	"io"
	"os"
)

// keep track of multiple writers (for adding and replacing)
var writers map[string]io.Writer = map[string]io.Writer{
	"default": os.Stdout,
}

// Add or replace a writer
// if name is nil the default writer will be replaced
func SetWriter(_writer io.Writer, name *string) {
	if name != nil {
		writers[*name] = _writer
	} else {
		writers["default"] = _writer
	}
	updateLogger()
}

// Remove a writer
func RemoveWriter(name string) error {
	if _, ok := writers[name]; ok {
		delete(writers, name)
		updateLogger()
		return nil
	} else {
		return fmt.Errorf("writer with name %s does not exist", name)
	}
}

// Internal function to set the loggers output according to all registered writers
func updateLogger() {
	allWriters := []io.Writer{}
	for _, writer := range writers {
		allWriters = append(allWriters, writer)
	}
	logger.SetOutput(io.MultiWriter(allWriters...))
}
