package ulog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// Helper function to select correct package, file, line and function for a logLine
func meta() *LogEntry {
	skip := 3
	var pc uintptr
	var functionRaw, fileRaw string
	var line int
	for {
	start:
		pc, fileRaw, line, _ = runtime.Caller(skip)
		functionRaw = runtime.FuncForPC(pc).Name()
		for _, skipFunction := range skipFunctions {
			if functionRaw == skipFunction {
				if globalDebug {
					fmt.Printf("skipping fn %s\n", functionRaw)
				}
				skip++
				// "just-in-case-failsafe"
				if skip > 100000 {
					break
				}
				goto start
			}
		}
		break

	}

	file := path.Base(fileRaw)
	packageEnd := strings.LastIndex(functionRaw, ".")
	packageName := functionRaw[:packageEnd]
	var functionName string
	if globalDebug {
		functionName = functionRaw
	} else {
		functionName = functionRaw[packageEnd+1:]
	}

	// Replace logic
	if val, ok := replaceFunctions[functionRaw]; ok {
		if globalDebug {
			fmt.Printf("replacing fn %s with %s\n", functionRaw, val)
		}
		functionName = ""
		packageName = val
		file = ""
		line = 0
	}

	return &LogEntry{
		Package:  packageName,
		File:     file,
		Line:     line,
		Function: functionName,
	}
}
