package v2

import (
	"go.uber.org/zap"
)

func L() *zap.Logger {
	return zap.L()
}

func S() *zap.SugaredLogger {
	return zap.S()
}

func Trace(v ...interface{}) {
	skipOneLogger.Warn("deprecated: TRACE")
	skipOneLogger.Debug(v...)
}
func Tracef(template string, v ...interface{}) {
	skipOneLogger.Warn("deprecated: TRACE")
	skipOneLogger.Debugf(template, v...)
}
func Debug(v ...interface{})                    { skipOneLogger.Debug(v...) }
func Debugf(template string, v ...interface{})  { skipOneLogger.Debugf(template, v...) }
func Info(v ...interface{})                     { skipOneLogger.Info(v...) }
func Infof(template string, v ...interface{})   { skipOneLogger.Infof(template, v...) }
func Warn(v ...interface{})                     { skipOneLogger.Warn(v...) }
func Warnf(template string, v ...interface{})   { skipOneLogger.Warnf(template, v...) }
func Error(v ...interface{})                    { skipOneLogger.Error(v...) }
func Errorf(template string, v ...interface{})  { skipOneLogger.Errorf(template, v...) }
func Fatal(v ...interface{})                    { skipOneLogger.Fatal(v...) }
func Fatalf(template string, v ...interface{})  { skipOneLogger.Fatalf(template, v...) }
func DPanic(v ...interface{})                   { skipOneLogger.DPanic(v...) }
func DPanicf(template string, v ...interface{}) { skipOneLogger.DPanicf(template, v...) }
func Panic(v ...interface{})                    { skipOneLogger.Panic(v...) }
func Panicf(template string, v ...interface{})  { skipOneLogger.Panicf(template, v...) }
