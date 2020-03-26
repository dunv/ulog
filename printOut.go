package ulog

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

// This is where all formatted log-output is actually printed
func printOut(fmtString *string, level logLevelString, v ...interface{}) {
	logLine := meta()
	logLine.Time = time.Now().Format(globalTsFormat) // TODO: customLogger support
	logLine.Level = level
	if fmtString != nil {
		logLine.Message = fmt.Sprintf(*fmtString, v...)
	} else {
		logLine.Message = fmt.Sprint(v...)
	}
	var tpl bytes.Buffer
	err := globalLineTemplate.Execute(&tpl, logLine) // TODO: customLogger support
	if err != nil {
		log.Printf("could not create ulogger line: %s", err)
	}
	logger.Print(tpl.String())
}
