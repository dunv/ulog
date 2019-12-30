package ulog

import (
	"log"
	"strings"
)

// A writer (implementing the io.Writer interface) which logs everything to its configured logger and to ERROR-level
// If not logger is configured, the default logger will be used
type UErrorWriter struct {
	Logger *ULogger
}

func (l UErrorWriter) Write(p []byte) (n int, err error) {
	cleanedString := strings.ReplaceAll(string(p), "\n", "")
	if l.Logger != nil {
		(*l.Logger).Error(cleanedString)
	} else {
		Error(cleanedString)
	}
	return len(p), nil
}

// Helper for instantiating UErrorWriter
func NewErrorLogger(ulogger *ULogger) *log.Logger {
	errorWriter := UErrorWriter{
		Logger: ulogger,
	}
	return log.New(errorWriter, "", 0)
}
