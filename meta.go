package ulog

import (
	"path"
	"runtime"
	"strings"
)

func meta() *LogEntry {
	skip := 3
	var pc uintptr
	var functionRaw, fileRaw string
	var line int
	for {
	start:
		pc, fileRaw, line, _ = runtime.Caller(skip)
		functionRaw = runtime.FuncForPC(pc).Name()
		// fmt.Println(functionRaw)
		for _, skipFunction := range skipFunctions {
			if functionRaw == skipFunction {
				skip++
				goto start
			}
		}
		break

	}
	file := path.Base(fileRaw)
	packageEnd := strings.LastIndex(functionRaw, ".")
	packageName := functionRaw[:packageEnd]
	var functionName string
	if debug {
		functionName = functionRaw
	} else {
		functionName = functionRaw[packageEnd+1:]
	}

	return &LogEntry{
		Package:  packageName,
		File:     file,
		Line:     line,
		Function: functionName,
	}
}
