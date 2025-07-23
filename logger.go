package ulog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// convenience for passing the logger to third-party libs
// wip, methods might need to be added
type Logger struct {
	log *zap.SugaredLogger
}

func NewDefaultLogger() Logger {
	return Logger{
		log: SkipOneS(),
	}
}

func NewLogger(log *zap.SugaredLogger) Logger {
	return Logger{
		log: log,
	}
}

func (l Logger) Fatalln(v ...any)               { l.log.Fatal(v...) }
func (l Logger) Panicln(v ...any)               { l.log.Panic(v...) }
func (l Logger) Print(v ...any)                 { l.log.Info(v...) }
func (l Logger) Printf(format string, v ...any) { l.log.Infof(format, v...) }
func (l Logger) Println(v ...any)               { l.log.Info(v...) }

func (l Logger) Debug(v ...any)                    { l.log.Debug(v...) }
func (l Logger) Debugf(template string, v ...any)  { l.log.Debugf(template, v...) }
func (l Logger) Info(v ...any)                     { l.log.Info(v...) }
func (l Logger) Infof(template string, v ...any)   { l.log.Infof(template, v...) }
func (l Logger) Warn(v ...any)                     { l.log.Warn(v...) }
func (l Logger) Warnf(template string, v ...any)   { l.log.Warnf(template, v...) }
func (l Logger) Error(v ...any)                    { l.log.Error(v...) }
func (l Logger) Errorf(template string, v ...any)  { l.log.Errorf(template, v...) }
func (l Logger) Fatal(v ...any)                    { l.log.Fatal(v...) }
func (l Logger) Fatalf(template string, v ...any)  { l.log.Fatalf(template, v...) }
func (l Logger) DPanic(v ...any)                   { l.log.DPanic(v...) }
func (l Logger) DPanicf(template string, v ...any) { l.log.DPanicf(template, v...) }
func (l Logger) Panic(v ...any)                    { l.log.Panic(v...) }
func (l Logger) Panicf(template string, v ...any)  { l.log.Panicf(template, v...) }

func (l Logger) Log(lvl zapcore.Level, args ...any) { l.log.Log(lvl, args...) }
func (l Logger) Logf(lvl zapcore.Level, template string, args ...any) {
	l.log.Logf(lvl, template, args...)
}
func (l Logger) Logw(lvl zapcore.Level, msg string, keysAndValues ...any) {
	l.log.Logw(lvl, msg, keysAndValues...)
}
