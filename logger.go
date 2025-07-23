package ulog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// convenience for passing the logger to third-party libs
type Logger interface {
	Fatalln(v ...any)
	Panicln(v ...any)
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
	Debug(v ...any)
	Debugf(template string, v ...any)
	Debugw(msg string, keysAndValues ...any)
	Info(v ...any)
	Infof(template string, v ...any)
	Infow(msg string, keysAndValues ...any)
	Warn(v ...any)
	Warnf(template string, v ...any)
	Warnw(msg string, keysAndValues ...any)
	Error(v ...any)
	Errorf(template string, v ...any)
	Errorw(msg string, keysAndValues ...any)
	Fatal(v ...any)
	Fatalf(template string, v ...any)
	Fatalw(msg string, keysAndValues ...any)
	DPanic(v ...any)
	DPanicf(template string, v ...any)
	DPanicw(msg string, keysAndValues ...any)
	Panic(v ...any)
	Panicf(template string, v ...any)
	Panicw(msg string, keysAndValues ...any)
	Log(lvl zapcore.Level, args ...any)
	Logf(lvl zapcore.Level, template string, args ...any)
	Logw(lvl zapcore.Level, msg string, keysAndValues ...any)
}

// convenience for passing the logger to third-party libs
type logger struct {
	log *zap.SugaredLogger
}

func NewDefaultLogger() Logger {
	return logger{
		log: SkipOneS(),
	}
}

func NewDiscardLogger() Logger {
	return logger{
		log: zap.NewNop().Sugar(),
	}
}

func NewLogger(log *zap.SugaredLogger) Logger {
	return logger{
		log: log,
	}
}

func (l logger) Fatalln(v ...any)                         { l.log.Fatal(v...) }
func (l logger) Panicln(v ...any)                         { l.log.Panic(v...) }
func (l logger) Print(v ...any)                           { l.log.Info(v...) }
func (l logger) Printf(format string, v ...any)           { l.log.Infof(format, v...) }
func (l logger) Println(v ...any)                         { l.log.Info(v...) }
func (l logger) Debug(v ...any)                           { l.log.Debug(v...) }
func (l logger) Debugf(template string, v ...any)         { l.log.Debugf(template, v...) }
func (l logger) Debugw(msg string, keysAndValues ...any)  { l.log.Debugw(msg, keysAndValues...) }
func (l logger) Info(v ...any)                            { l.log.Info(v...) }
func (l logger) Infof(template string, v ...any)          { l.log.Infof(template, v...) }
func (l logger) Infow(msg string, keysAndValues ...any)   { l.log.Infow(msg, keysAndValues...) }
func (l logger) Warn(v ...any)                            { l.log.Warn(v...) }
func (l logger) Warnf(template string, v ...any)          { l.log.Warnf(template, v...) }
func (l logger) Warnw(msg string, keysAndValues ...any)   { l.log.Warnw(msg, keysAndValues...) }
func (l logger) Error(v ...any)                           { l.log.Error(v...) }
func (l logger) Errorf(template string, v ...any)         { l.log.Errorf(template, v...) }
func (l logger) Errorw(msg string, keysAndValues ...any)  { l.log.Errorw(msg, keysAndValues...) }
func (l logger) Fatal(v ...any)                           { l.log.Fatal(v...) }
func (l logger) Fatalf(template string, v ...any)         { l.log.Fatalf(template, v...) }
func (l logger) Fatalw(msg string, keysAndValues ...any)  { l.log.Fatalw(msg, keysAndValues...) }
func (l logger) DPanic(v ...any)                          { l.log.DPanic(v...) }
func (l logger) DPanicf(template string, v ...any)        { l.log.DPanicf(template, v...) }
func (l logger) DPanicw(msg string, keysAndValues ...any) { l.log.DPanicw(msg, keysAndValues...) }
func (l logger) Panic(v ...any)                           { l.log.Panic(v...) }
func (l logger) Panicf(template string, v ...any)         { l.log.Panicf(template, v...) }
func (l logger) Panicw(msg string, keysAndValues ...any)  { l.log.Panicw(msg, keysAndValues...) }

func (l logger) Log(lvl zapcore.Level, args ...any) { l.log.Log(lvl, args...) }
func (l logger) Logf(lvl zapcore.Level, template string, args ...any) {
	l.log.Logf(lvl, template, args...)
}
func (l logger) Logw(lvl zapcore.Level, msg string, keysAndValues ...any) {
	l.log.Logw(lvl, msg, keysAndValues...)
}
