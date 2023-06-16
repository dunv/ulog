package v2

import "go.uber.org/zap"

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

func (l Logger) Fatal(v ...interface{}) {
	l.log.Fatal(v...)
}

func (l Logger) Fatalf(format string, v ...interface{}) {
	l.log.Fatalf(format, v...)
}

func (l Logger) Fatalln(v ...interface{}) {
	l.log.Fatal(v...)
}

func (l Logger) Panic(v ...interface{}) {
	l.log.Panic(v...)
}

func (l Logger) Panicf(format string, v ...interface{}) {
	l.log.Panicf(format, v...)
}

func (l Logger) Panicln(v ...interface{}) {
	l.log.Panic(v...)
}

func (l Logger) Print(v ...interface{}) {
	l.log.Info(v...)
}

func (l Logger) Printf(format string, v ...interface{}) {
	l.log.Infof(format, v...)
}

func (l Logger) Println(v ...interface{}) {
	l.log.Info(v...)
}
