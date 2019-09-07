package ulog

import (
	"html/template"
	"io"
	"os"
)

var writer io.Writer
var lineTemplate *template.Template
var tsFormat string
var skipFunctions []string
var logLevel LogLevel

func Init(_writer io.Writer, _level LogLevel, _skipFunctions ...string) {

	defaultSkipFunctions := []string{
		"github.com/dunv/ulog.LogIfError",
		"github.com/dunv/ulog.LogIfErrorSecondArg",
		"github.com/dunv/ulog.LogIfErrorToInfo",
		"github.com/dunv/ulog.LogIfErrorToInfoSecondArg",
	}

	writer = _writer
	fmtString := "{{ .Time }} | {{ .Level }} | n/a | {{ .Package }} {{ .File }}:{{ .Line }} ({{ .Function }}) | {{ .Message }}\n"
	tsFormat = "2006-01-02 15:04:05.000"
	lineTemplate = template.Must(template.New("lineTemplate").Parse(fmtString))
	skipFunctions = append(defaultSkipFunctions, _skipFunctions...)
	logLevel = _level
}

func Trace(v ...interface{}) {
	if logLevel <= LEVEL_TRACE {
		printOut(nil, _LEVEL_TRACE, v...)
	}
}

func Tracef(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_TRACE {
		printOut(&fmtString, _LEVEL_TRACE, v...)
	}
}

func Debug(v ...interface{}) {
	if logLevel <= LEVEL_DEBUG {
		printOut(nil, _LEVEL_DEBUG, v...)
	}

}
func Debugf(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_DEBUG {
		printOut(&fmtString, _LEVEL_DEBUG, v...)
	}
}

func Info(v ...interface{}) {
	if logLevel <= LEVEL_INFO {
		printOut(nil, _LEVEL_INFO, v...)
	}
}

func Infof(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_INFO {
		printOut(&fmtString, _LEVEL_INFO, v...)
	}
}

func Warn(v ...interface{}) {
	if logLevel <= LEVEL_WARNING {
		printOut(nil, _LEVEL_WARNING, v...)
	}
}

func Warnf(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_WARNING {
		printOut(&fmtString, _LEVEL_WARNING, v...)
	}
}

func Error(v ...interface{}) {
	if logLevel <= LEVEL_ERROR {
		printOut(nil, _LEVEL_ERROR, v...)
	}
}

func Errorf(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_ERROR {
		printOut(&fmtString, _LEVEL_ERROR, v...)
	}
}

func Panic(v ...interface{}) {
	if logLevel <= LEVEL_ERROR {
		printOut(nil, _LEVEL_ERROR, v...)
	}
	panic(v)
}

func Panicf(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_ERROR {
		printOut(&fmtString, _LEVEL_ERROR, v...)
	}
	panic(v)
}

func Fatal(v ...interface{}) {
	if logLevel <= LEVEL_FATAL {
		printOut(nil, _LEVEL_FATAL, v...)
	}
	os.Exit(1)
}

func Fatalf(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_FATAL {
		printOut(&fmtString, _LEVEL_FATAL, v...)
	}
	os.Exit(1)
}
