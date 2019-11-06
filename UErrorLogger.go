package ulog

import (
	"log"
	"strings"
)

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

func NewErrorLogger(ulogger *ULogger) *log.Logger {
	errorWriter := UErrorWriter{
		Logger: ulogger,
	}
	return log.New(errorWriter, "", 0)
}
