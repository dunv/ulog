package ulog

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
func Debug(v ...any)                    { skipOneSugaredLogger.Debug(v...) }
func Debugf(template string, v ...any)  { skipOneSugaredLogger.Debugf(template, v...) }
func Info(v ...any)                     { skipOneSugaredLogger.Info(v...) }
func Infof(template string, v ...any)   { skipOneSugaredLogger.Infof(template, v...) }
func Warn(v ...any)                     { skipOneSugaredLogger.Warn(v...) }
func Warnf(template string, v ...any)   { skipOneSugaredLogger.Warnf(template, v...) }
func Error(v ...any)                    { skipOneSugaredLogger.Error(v...) }
func Errorf(template string, v ...any)  { skipOneSugaredLogger.Errorf(template, v...) }
func Fatal(v ...any)                    { skipOneSugaredLogger.Fatal(v...) }
func Fatalf(template string, v ...any)  { skipOneSugaredLogger.Fatalf(template, v...) }
func DPanic(v ...any)                   { skipOneSugaredLogger.DPanic(v...) }
func DPanicf(template string, v ...any) { skipOneSugaredLogger.DPanicf(template, v...) }
func Panic(v ...any)                    { skipOneSugaredLogger.Panic(v...) }
func Panicf(template string, v ...any)  { skipOneSugaredLogger.Panicf(template, v...) }

func With(args ...any) *zap.SugaredLogger {
	return skipOneSugaredLogger.With(args...)
}
