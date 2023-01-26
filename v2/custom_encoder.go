package v2

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	color_json    string = "\033[90m"
	color_trace   string = "\033[90m"
	color_debug   string = "\033[34m"
	color_warning string = "\033[93m"
	color_error   string = "\033[91m"
	color_fatal   string = "\033[95m"
	color_end     string = "\033[0m"
)

var (
	levelToColorStart = map[zapcore.Level]string{
		zapcore.TraceLevel:  color_trace,
		zapcore.DebugLevel:  color_debug,
		zapcore.InfoLevel:   "",
		zapcore.WarnLevel:   color_warning,
		zapcore.ErrorLevel:  color_error,
		zapcore.DPanicLevel: color_fatal,
		zapcore.PanicLevel:  color_fatal,
		zapcore.FatalLevel:  color_fatal,
	}
	levelString     = make(map[zapcore.Level]string, len(levelToColorStart))
	levelFieldWidth int
)

func init() {
	for level := range levelToColorStart {
		lvlStr := level.CapitalString()
		levelString[level] = fmt.Sprintf("%-6s", lvlStr)
		if levelFieldWidth < len(lvlStr) {
			levelFieldWidth = len(lvlStr)
		}
	}
}

type customEncoder struct {
	zapcore.Encoder
	separator  string
	bufferpool buffer.Pool
}

func newCustomEncoder() zapcore.Encoder {
	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "loggerName",
		CallerKey:      "location",
		FunctionKey:    "function",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return customEncoder{
		Encoder:    zapcore.NewJSONEncoder(cfg),
		separator:  " | ",
		bufferpool: buffer.NewPool(),
	}
}

func (c customEncoder) Clone() zapcore.Encoder {
	return customEncoder{Encoder: c.Encoder.Clone()}
}

func (c customEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	line := c.bufferpool.Get()

	// Coloring
	if s, ok := levelToColorStart[ent.Level]; ok {
		line.AppendString(s)
	}
	line.AppendString(ent.Time.Format(time.RFC3339))
	line.AppendString(c.separator)
	appendPaddedLevel(ent.Level, line)
	line.AppendString(c.separator)
	appendCaller(ent.Caller, line)
	line.AppendString(c.separator)
	line.AppendString(ent.Message)

	if !selectedOptions.stripAdditionalFields {
		line.AppendString(c.separator)
		line.AppendString(color_json)
		buf, _ := c.Encoder.EncodeEntry(ent, fields)
		_, _ = line.Write(buf.Bytes()[:len(buf.Bytes())-1])
		buf.Free()
		line.AppendString(color_end)
	}

	if ent.Stack != "" {
		line.AppendByte('\n')
		line.AppendString(ent.Stack)
	}

	// Coloring
	if _, ok := levelToColorStart[ent.Level]; ok {
		line.AppendString(color_end)
	}
	line.AppendString(zapcore.DefaultLineEnding)

	return line, nil
}

var appendPaddedLevel = func(l zapcore.Level, enc *buffer.Buffer) {
	if s, ok := levelString[l]; ok {
		enc.AppendString(s)
	}
}

// copied from zapcore with constant width
var appendCaller = func(caller zapcore.EntryCaller, enc *buffer.Buffer) {
	if !caller.Defined {
		enc.AppendString("caller undefined")
		return
	}
	// Find the last separator.
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
	str := fmt.Sprintf("%s:%d", caller.File[idx+1:], caller.Line)
	// Pad or cut
	if selectedOptions.callerFieldWidth != -1 {
		if len(str) > selectedOptions.callerFieldWidth {
			str = str[len(str)-selectedOptions.callerFieldWidth:]
		} else {
			str = str + strings.Repeat(" ", selectedOptions.callerFieldWidth-len(str))
		}
	}
	enc.AppendString(str)
}
