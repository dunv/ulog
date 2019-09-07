package ulog

import (
	"fmt"
	"html/template"
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
	err := template.Must(lineTemplate.Clone()).Execute(writer, logLine)
	if err != nil {
		log.Printf("could not create ulogger line: %s", err)
	}

}
