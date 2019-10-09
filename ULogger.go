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
	LogIfError(err error)
	LogIfErrorSecondArg(tmp interface{}, err error)
	PanicIfError(err error)
	PanicIfErrorSecondArg(tmp interface{}, err error)
	LogIfErrorToInfo(err error)
	LogIfErrorToInfoSecondArg(tmp interface{}, err error)
}

type ULog struct{}

func (l ULog) Trace(v ...interface{})                           { Trace(v...) }
func (l ULog) Tracef(fmtString string, v ...interface{})        { Tracef(fmtString, v...) }
func (l ULog) Debug(v ...interface{})                           { Debug(v...) }
func (l ULog) Debugf(fmtString string, v ...interface{})        { Debugf(fmtString, v...) }
func (l ULog) Info(v ...interface{})                            { Info(v...) }
func (l ULog) Infof(fmtString string, v ...interface{})         { Infof(fmtString, v...) }
func (l ULog) Warn(v ...interface{})                            { Warn(v...) }
func (l ULog) Warnf(fmtString string, v ...interface{})         { Warnf(fmtString, v...) }
func (l ULog) Error(v ...interface{})                           { Error(v...) }
func (l ULog) Errorf(fmtString string, v ...interface{})        { Errorf(fmtString, v...) }
func (l ULog) Panic(v ...interface{})                           { Panic(v...) }
func (l ULog) Panicf(fmtString string, v ...interface{})        { Panicf(fmtString, v...) }
func (l ULog) Fatal(v ...interface{})                           { Fatal(v...) }
func (l ULog) Fatalf(fmtString string, v ...interface{})        { Fatalf(fmtString, v...) }
func (l ULog) LogIfError(err error)                             { LogIfError(err) }
func (l ULog) LogIfErrorSecondArg(tmp interface{}, err error)   { LogIfErrorSecondArg(tmp, err) }
func (l ULog) PanicIfError(err error)                           { PanicIfError(err) }
func (l ULog) PanicIfErrorSecondArg(tmp interface{}, err error) { PanicIfErrorSecondArg(tmp, err) }
func (l ULog) LogIfErrorToInfo(err error)                       { LogIfErrorToInfo(err) }
func (l ULog) LogIfErrorToInfoSecondArg(tmp interface{}, err error) {
	LogIfErrorToInfoSecondArg(tmp, err)
}

func NewUlog() *ULog {
	return &ULog{}
}
