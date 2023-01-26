package v2

type ULogger interface {
	Trace(v ...interface{})
	Tracef(template string, v ...interface{})
	Debug(v ...interface{})
	Debugf(template string, v ...interface{})
	Info(v ...interface{})
	Infof(template string, v ...interface{})
	Warn(v ...interface{})
	Warnf(template string, v ...interface{})
	Error(v ...interface{})
	Errorf(template string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(template string, v ...interface{})
	DPanic(v ...interface{})
	DPanicf(template string, v ...interface{})
	Panic(v ...interface{})
	Panicf(template string, v ...interface{})
}

type Logger struct{}

func (l Logger) Trace(v ...interface{}) {
	skipOneLogger.Warn("deprecated: TRACE")
	skipOneLogger.Debug(v...)
}
func (l Logger) Tracef(template string, v ...interface{}) {
	skipOneLogger.Warn("deprecated: TRACE")
	skipOneLogger.Debugf(template, v...)
}
func (l Logger) Debug(v ...interface{})                    { skipOneLogger.Debug(v...) }
func (l Logger) Debugf(template string, v ...interface{})  { skipOneLogger.Debugf(template, v...) }
func (l Logger) Info(v ...interface{})                     { skipOneLogger.Info(v...) }
func (l Logger) Infof(template string, v ...interface{})   { skipOneLogger.Infof(template, v...) }
func (l Logger) Warn(v ...interface{})                     { skipOneLogger.Warn(v...) }
func (l Logger) Warnf(template string, v ...interface{})   { skipOneLogger.Warnf(template, v...) }
func (l Logger) Error(v ...interface{})                    { skipOneLogger.Error(v...) }
func (l Logger) Errorf(template string, v ...interface{})  { skipOneLogger.Errorf(template, v...) }
func (l Logger) Fatal(v ...interface{})                    { skipOneLogger.Fatal(v...) }
func (l Logger) Fatalf(template string, v ...interface{})  { skipOneLogger.Fatalf(template, v...) }
func (l Logger) DPanic(v ...interface{})                   { skipOneLogger.DPanic(v...) }
func (l Logger) DPanicf(template string, v ...interface{}) { skipOneLogger.DPanicf(template, v...) }
func (l Logger) Panic(v ...interface{})                    { skipOneLogger.Panic(v...) }
func (l Logger) Panicf(template string, v ...interface{})  { skipOneLogger.Panicf(template, v...) }
