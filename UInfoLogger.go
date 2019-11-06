package ulog

import (
	"log"
	"strings"
)

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

func NewInfoLogger(ulogger *ULogger) *log.Logger {
	infoWriter := UInfoWriter{
		Logger: ulogger,
	}
	return log.New(infoWriter, "", 0)
}
