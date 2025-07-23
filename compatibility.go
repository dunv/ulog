package ulog

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_globalMu sync.RWMutex
	_globalL  = zap.L()
	_globalS  = zap.S()
)

func SkipOneS() *zap.SugaredLogger {
	_globalMu.RLock()
	s := _globalS
	_globalMu.RUnlock()
	return s
}

func SkipOneL() *zap.Logger {
	_globalMu.RLock()
	l := _globalL
	_globalMu.RUnlock()
	return l
}

func L() *zap.Logger {
	return zap.L()
}

func S() *zap.SugaredLogger {
	return zap.S()
}

func Log(lvl zapcore.Level, args ...any) { SkipOneS().Log(lvl, args...) }
func Logf(lvl zapcore.Level, template string, args ...any) {
	SkipOneS().Logf(lvl, template, args...)
}
func Logw(lvl zapcore.Level, msg string, keysAndValues ...any) {
	SkipOneS().Logw(lvl, msg, keysAndValues...)
}
func Debug(v ...any)                           { SkipOneS().Debug(v...) }
func Debugf(template string, v ...any)         { SkipOneS().Debugf(template, v...) }
func Debugw(msg string, keysAndValues ...any)  { SkipOneS().Debugw(msg, keysAndValues...) }
func Info(v ...any)                            { SkipOneS().Info(v...) }
func Infof(template string, v ...any)          { SkipOneS().Infof(template, v...) }
func Infofw(msg string, keysAndValues ...any)  { SkipOneS().Infow(msg, keysAndValues...) }
func Warn(v ...any)                            { SkipOneS().Warn(v...) }
func Warnf(template string, v ...any)          { SkipOneS().Warnf(template, v...) }
func Warnw(msg string, keysAndValues ...any)   { SkipOneS().Warnw(msg, keysAndValues...) }
func Error(v ...any)                           { SkipOneS().Error(v...) }
func Errorf(template string, v ...any)         { SkipOneS().Errorf(template, v...) }
func Errorw(msg string, keysAndValues ...any)  { SkipOneS().Errorw(msg, keysAndValues...) }
func Fatal(v ...any)                           { SkipOneS().Fatal(v...) }
func Fatalf(template string, v ...any)         { SkipOneS().Fatalf(template, v...) }
func Fatalw(msg string, keysAndValues ...any)  { SkipOneS().Fatalw(msg, keysAndValues...) }
func DPanic(v ...any)                          { SkipOneS().DPanic(v...) }
func DPanicf(template string, v ...any)        { SkipOneS().DPanicf(template, v...) }
func Dpanicw(msg string, keysAndValues ...any) { SkipOneS().DPanicw(msg, keysAndValues...) }
func Panic(v ...any)                           { SkipOneS().Panic(v...) }
func Panicf(template string, v ...any)         { SkipOneS().Panicf(template, v...) }
func Panicw(msg string, keysAndValues ...any)  { SkipOneS().Panicw(msg, keysAndValues...) }

func With(args ...any) *zap.SugaredLogger {
	return S().With(args...)
}
