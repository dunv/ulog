package v2

import (
	"go.uber.org/zap"
)

// convenience for passing the logger to third-party libs
// wip, methods might need to be added
type Logger struct {
	log *zap.SugaredLogger
}

func NewDefaultLogger() Logger {
	return Logger{
		log: skipOneSugaredLogger,
	}
}

func NewLogger(log *zap.SugaredLogger) Logger {
	return Logger{
		log: log,
	}
}

func (l Logger) Fatalln(v ...interface{})               { l.log.Fatal(v...) }
func (l Logger) Panicln(v ...interface{})               { l.log.Panic(v...) }
func (l Logger) Print(v ...interface{})                 { l.log.Info(v...) }
func (l Logger) Printf(format string, v ...interface{}) { l.log.Infof(format, v...) }
func (l Logger) Println(v ...interface{})               { l.log.Info(v...) }

func (l Logger) Trace(v ...interface{})                    { l.log.Trace(v...) }
func (l Logger) Tracef(template string, v ...interface{})  { l.log.Tracef(template, v...) }
func (l Logger) Debug(v ...interface{})                    { l.log.Debug(v...) }
func (l Logger) Debugf(template string, v ...interface{})  { l.log.Debugf(template, v...) }
func (l Logger) Info(v ...interface{})                     { l.log.Info(v...) }
func (l Logger) Infof(template string, v ...interface{})   { l.log.Infof(template, v...) }
func (l Logger) Warn(v ...interface{})                     { l.log.Warn(v...) }
func (l Logger) Warnf(template string, v ...interface{})   { l.log.Warnf(template, v...) }
func (l Logger) Error(v ...interface{})                    { l.log.Error(v...) }
func (l Logger) Errorf(template string, v ...interface{})  { l.log.Errorf(template, v...) }
func (l Logger) Fatal(v ...interface{})                    { l.log.Fatal(v...) }
func (l Logger) Fatalf(template string, v ...interface{})  { l.log.Fatalf(template, v...) }
func (l Logger) DPanic(v ...interface{})                   { l.log.DPanic(v...) }
func (l Logger) DPanicf(template string, v ...interface{}) { l.log.DPanicf(template, v...) }
func (l Logger) Panic(v ...interface{})                    { l.log.Panic(v...) }
func (l Logger) Panicf(template string, v ...interface{})  { l.log.Panicf(template, v...) }
