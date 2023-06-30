package ulog

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

// Deprecated: use v2
func Trace(v ...interface{}) {
	if GetLogLevel() <= LEVEL_TRACE {
		printOut(nil, _LEVEL_TRACE, v...)
	}
}

// Deprecated: use v2
func Tracef(fmtString string, v ...interface{}) {
	if GetLogLevel() <= LEVEL_TRACE {
		printOut(&fmtString, _LEVEL_TRACE, v...)
	}
}

// Deprecated: use v2
func Debug(v ...interface{}) {
	if GetLogLevel() <= LEVEL_DEBUG {
		printOut(nil, _LEVEL_DEBUG, v...)
	}

}

// Deprecated: use v2
func Debugf(fmtString string, v ...interface{}) {
	if GetLogLevel() <= LEVEL_DEBUG {
		printOut(&fmtString, _LEVEL_DEBUG, v...)
	}
}

// Deprecated: use v2
func Info(v ...interface{}) {
	if GetLogLevel() <= LEVEL_INFO {
		printOut(nil, _LEVEL_INFO, v...)
	}
}

// Deprecated: use v2
func Infof(fmtString string, v ...interface{}) {
	if GetLogLevel() <= LEVEL_INFO {
		printOut(&fmtString, _LEVEL_INFO, v...)
	}
}

// Deprecated: use v2
func Warn(v ...interface{}) {
	if GetLogLevel() <= LEVEL_WARNING {
		printOut(nil, _LEVEL_WARNING, v...)
	}
}

// Deprecated: use v2
func Warnf(fmtString string, v ...interface{}) {
	if GetLogLevel() <= LEVEL_WARNING {
		printOut(&fmtString, _LEVEL_WARNING, v...)
	}
}

// Deprecated: use v2
func Error(v ...interface{}) {
	if GetLogLevel() <= LEVEL_ERROR {
		printOut(nil, _LEVEL_ERROR, v...)
	}
}

// Deprecated: use v2
func Errorf(fmtString string, v ...interface{}) {
	if GetLogLevel() <= LEVEL_ERROR {
		printOut(&fmtString, _LEVEL_ERROR, v...)
	}
}

// Deprecated: use v2
func Fatal(v ...interface{}) {
	if GetLogLevel() <= LEVEL_FATAL {
		printOut(nil, _LEVEL_FATAL, v...)
	}
	// This way we can make it testable
	if !globalDebug {
		os.Exit(1)
	} else {
		printOut(nil, _LEVEL_FATAL, "would os.Exit(1), but debug is enabled, NOT exiting.")
	}
}

// Deprecated: use v2
func Fatalf(fmtString string, v ...interface{}) {
	if GetLogLevel() <= LEVEL_FATAL {
		printOut(&fmtString, _LEVEL_FATAL, v...)
	}
	// This way we can make it testable
	if !globalDebug {
		os.Exit(1)
	} else {
		printOut(nil, _LEVEL_FATAL, "would os.Exit(1), but debug is enabled, NOT exiting.")
	}
}

// Deprecated: use v2
func Panic(v interface{}) {
	if GetLogLevel() <= LEVEL_ERROR {
		printOut(nil, _LEVEL_ERROR, v)
	}
	panic(v)
}

// Deprecated: use v2
func Panicf(fmtString string, v ...interface{}) {
	if GetLogLevel() <= LEVEL_ERROR {
		printOut(&fmtString, _LEVEL_ERROR, v...)
	}
	panic(v)
}

// Deprecated: use v2
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

// Deprecated: use v2
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
