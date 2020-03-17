package ulog

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

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

func Fatal(v ...interface{}) {
	if logLevel <= LEVEL_FATAL {
		printOut(nil, _LEVEL_FATAL, v...)
	}
	// This way we can make it testable
	if !debug {
		printOut(nil, _LEVEL_FATAL, "would os.Exit(1), but debug is enabled, NOT exiting.")
		os.Exit(1)
	}
}

func Fatalf(fmtString string, v ...interface{}) {
	if logLevel <= LEVEL_FATAL {
		printOut(&fmtString, _LEVEL_FATAL, v...)
	}
	// This way we can make it testable
	if !debug {
		printOut(nil, _LEVEL_FATAL, "would os.Exit(1), but debug is enabled, NOT exiting.")
		os.Exit(1)
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

func LogWithLevel(level LogLevel, v ...interface{}) {
	switch level {
	case LEVEL_TRACE:
		Trace(v...)
	case LEVEL_DEBUG:
		Debug(v...)
	case LEVEL_INFO:
		Info(v...)
	case LEVEL_WARNING:
		Warn(v...)
	case LEVEL_ERROR:
		Error(v...)
	case LEVEL_FATAL:
		Fatal(v...)
	}
}

func LogWithLevelf(level LogLevel, fmtString string, v ...interface{}) {
	switch level {
	case LEVEL_TRACE:
		Tracef(fmtString, v...)
	case LEVEL_DEBUG:
		Debugf(fmtString, v...)
	case LEVEL_INFO:
		Infof(fmtString, v...)
	case LEVEL_WARNING:
		Warnf(fmtString, v...)
	case LEVEL_ERROR:
		Errorf(fmtString, v...)
	case LEVEL_FATAL:
		Fatalf(fmtString, v...)
	}
}
