package ulog

type ULogger interface {
	Trace(v ...interface{})
	Tracef(fmtString string, v ...interface{})
	Debug(v ...interface{})
	Debugf(fmtString string, v ...interface{})
	Info(v ...interface{})
	Infof(fmtString string, v ...interface{})
	Warn(v ...interface{})
	Warnf(fmtString string, v ...interface{})
	Error(v ...interface{})
	Errorf(fmtString string, v ...interface{})
	Panic(v ...interface{})
	Panicf(fmtString string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(fmtString string, v ...interface{})
}
