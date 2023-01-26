package v2

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

var consoleLevelEncoder = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString(fmt.Sprintf("%sDEBUG %s", COLOR_DEBUG, COLOR_END))
	case zapcore.InfoLevel:
		enc.AppendString(fmt.Sprintf("%sINFO  %s", COLOR_INFO, COLOR_END))
	case zapcore.WarnLevel:
		enc.AppendString(fmt.Sprintf("%sWARN  %s", COLOR_WARNING, COLOR_END))
	case zapcore.ErrorLevel:
		enc.AppendString(fmt.Sprintf("%sERROR %s", COLOR_ERROR, COLOR_END))
	case zapcore.DPanicLevel:
		enc.AppendString(fmt.Sprintf("%sPANIC %s", COLOR_FATAL, COLOR_END))
	case zapcore.PanicLevel:
		enc.AppendString(fmt.Sprintf("%sPANIC %s", COLOR_FATAL, COLOR_END))
	case zapcore.FatalLevel:
		enc.AppendString(fmt.Sprintf("%sFATAL %s", COLOR_FATAL, COLOR_END))
	default:
		enc.AppendString(fmt.Sprintf("LEVEL(%d)", l))
	}
}

var consoleTimeEncoder = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

var consoleCallerEncoder = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// copied from source
	if !caller.Defined {
		enc.AppendString("undefined")
		return
	}
	// nb. To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	//
	// Find the last separator.
	//
	idx := strings.LastIndexByte(caller.File, '/')
	if idx == -1 {
		enc.AppendString(caller.FullPath())
		return
	}
	// Find the penultimate separator.
	idx = strings.LastIndexByte(caller.File[:idx], '/')
	if idx == -1 {
		enc.AppendString(caller.FullPath())
		return
	}
	// Keep everything after the penultimate separator.
	enc.AppendString(fmt.Sprintf("%s:%-4d", caller.File[idx+1:], caller.Line))
}

var consoleEncoderConfig = zapcore.EncoderConfig{
	TimeKey:          "T",
	LevelKey:         "L",
	NameKey:          "N",
	CallerKey:        "C",
	FunctionKey:      zapcore.OmitKey,
	MessageKey:       "M",
	StacktraceKey:    "S",
	LineEnding:       zapcore.DefaultLineEnding,
	ConsoleSeparator: " | ",
	EncodeLevel:      consoleLevelEncoder,
	EncodeTime:       consoleTimeEncoder,
	EncodeDuration:   zapcore.StringDurationEncoder,
	EncodeCaller:     consoleCallerEncoder,
}
