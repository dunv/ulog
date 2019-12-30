package ulog

import (
	"log"
	"strings"
)

// A writer (implementing the io.Writer interface) which logs everything to its configured logger and to INFO-level
// If not logger is configured, the default logger will be used
type UInfoWriter struct {
	Logger *ULogger
}

func (l UInfoWriter) Write(p []byte) (n int, err error) {
	cleanedString := strings.ReplaceAll(string(p), "\n", "")
	if l.Logger != nil {
		(*l.Logger).Info(cleanedString)
	} else {
		Info(cleanedString)
	}
	return len(p), nil
}

// Helper for instantiating UInfoWriter
func NewInfoLogger(ulogger *ULogger) *log.Logger {
	infoWriter := UInfoWriter{
		Logger: ulogger,
	}
	return log.New(infoWriter, "", 0)
}
