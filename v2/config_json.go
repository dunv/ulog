package v2

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

var jsonLevelEncoder = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString("DEBUG")
	case zapcore.InfoLevel:
		enc.AppendString("INFO")
	case zapcore.WarnLevel:
		enc.AppendString("WARN")
	case zapcore.ErrorLevel:
		enc.AppendString("ERROR")
	case zapcore.DPanicLevel:
		enc.AppendString("PANIC")
	case zapcore.PanicLevel:
		enc.AppendString("PANIC")
	case zapcore.FatalLevel:
		enc.AppendString("FATAL")
	default:
		enc.AppendString(fmt.Sprintf("LEVEL(%d)", l))
	}
}

var jsonTimeEncoder = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}

var jsonEncoderConfig = zapcore.EncoderConfig{
	TimeKey:          "time",
	LevelKey:         "level",
	NameKey:          "name",
	CallerKey:        "caller",
	FunctionKey:      "function",
	MessageKey:       "message",
	StacktraceKey:    "stackTrace",
	LineEnding:       zapcore.DefaultLineEnding,
	ConsoleSeparator: " | ",
	EncodeLevel:      jsonLevelEncoder,
	EncodeTime:       jsonTimeEncoder,
	EncodeDuration:   zapcore.StringDurationEncoder,
	EncodeCaller:     zapcore.ShortCallerEncoder,
}
