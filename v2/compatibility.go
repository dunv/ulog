package v2

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func L() *zap.Logger {
	return zap.L()
}

func SkipOneL() *zap.Logger {
	return skipOneLogger
}

func S() *zap.SugaredLogger {
	return zap.S()
}

func SkipOneS() *zap.SugaredLogger {
	return skipOneSugaredLogger
}

func Log(lvl zapcore.Level, msg string, fields ...zapcore.Field) {
	skipOneLogger.Log(lvl, msg, fields...)
}
func Trace(v ...interface{})                    { skipOneSugaredLogger.Trace(v...) }
func Tracef(template string, v ...interface{})  { skipOneSugaredLogger.Tracef(template, v...) }
func Debug(v ...interface{})                    { skipOneSugaredLogger.Debug(v...) }
func Debugf(template string, v ...interface{})  { skipOneSugaredLogger.Debugf(template, v...) }
func Info(v ...interface{})                     { skipOneSugaredLogger.Info(v...) }
func Infof(template string, v ...interface{})   { skipOneSugaredLogger.Infof(template, v...) }
func Warn(v ...interface{})                     { skipOneSugaredLogger.Warn(v...) }
func Warnf(template string, v ...interface{})   { skipOneSugaredLogger.Warnf(template, v...) }
func Error(v ...interface{})                    { skipOneSugaredLogger.Error(v...) }
func Errorf(template string, v ...interface{})  { skipOneSugaredLogger.Errorf(template, v...) }
func Fatal(v ...interface{})                    { skipOneSugaredLogger.Fatal(v...) }
func Fatalf(template string, v ...interface{})  { skipOneSugaredLogger.Fatalf(template, v...) }
func DPanic(v ...interface{})                   { skipOneSugaredLogger.DPanic(v...) }
func DPanicf(template string, v ...interface{}) { skipOneSugaredLogger.DPanicf(template, v...) }
func Panic(v ...interface{})                    { skipOneSugaredLogger.Panic(v...) }
func Panicf(template string, v ...interface{})  { skipOneSugaredLogger.Panicf(template, v...) }
