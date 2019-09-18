package ulog

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

func printOut(fmtString *string, level logLevelString, v ...interface{}) {
	logLine := meta()
	logLine.Time = time.Now().Format(tsFormat)
	logLine.Level = level
	if fmtString != nil {
		logLine.Message = fmt.Sprintf(*fmtString, v...)
	} else {
		logLine.Message = fmt.Sprint(v...)
	}
	var tpl bytes.Buffer
	err := lineTemplate.Execute(&tpl, logLine)
	if err != nil {
		log.Printf("could not create ulogger line: %s", err)
	}
	logger.Printf(tpl.String())
}
