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
