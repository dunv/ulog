package v2

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	colorJson    string = "\033[90m"
	colorTrace   string = "\033[90m"
	colorDebug   string = "\033[34m"
	colorWarning string = "\033[93m"
	colorError   string = "\033[91m"
	colorFatal   string = "\033[95m"
	colorEnd     string = "\033[0m"
	timeFormat          = time.RFC3339
)

var (
	levelToColorStart = map[zapcore.Level][]byte{
		zapcore.TraceLevel:  []byte(colorTrace),
		zapcore.DebugLevel:  []byte(colorDebug),
		zapcore.InfoLevel:   nil,
		zapcore.WarnLevel:   []byte(colorWarning),
		zapcore.ErrorLevel:  []byte(colorError),
		zapcore.DPanicLevel: []byte(colorFatal),
		zapcore.PanicLevel:  []byte(colorFatal),
		zapcore.FatalLevel:  []byte(colorFatal),
	}
	levelToColorEnd = map[zapcore.Level][]byte{
		zapcore.TraceLevel:  []byte(colorEnd),
		zapcore.DebugLevel:  []byte(colorEnd),
		zapcore.InfoLevel:   nil,
		zapcore.WarnLevel:   []byte(colorEnd),
		zapcore.ErrorLevel:  []byte(colorEnd),
		zapcore.DPanicLevel: []byte(colorEnd),
		zapcore.PanicLevel:  []byte(colorEnd),
		zapcore.FatalLevel:  []byte(colorEnd),
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
		EncodeTime:     zapcore.TimeEncoderOfLayout(timeFormat),
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
	line.AppendBytes(levelToColorStart[ent.Level])
	line.AppendString(ent.Time.Format(timeFormat))
	line.AppendString(c.separator)
	appendPaddedLevel(ent.Level, line)
	line.AppendString(c.separator)
	if selectedOptions.renderDummyThread {
		line.AppendString("n/a")
		line.AppendString(c.separator)
	}
	appendCaller(ent.Caller, line)
	line.AppendString(c.separator)
	line.AppendString(ent.Message)

	if !selectedOptions.stripAdditionalFields {
		line.AppendString(c.separator)
		line.AppendString(colorJson)
		buf, _ := c.Encoder.EncodeEntry(ent, fields)
		_, _ = line.Write(buf.Bytes()[:len(buf.Bytes())-1])
		buf.Free()
		line.AppendString(colorEnd)
	}

	if ent.Stack != "" {
		line.AppendByte('\n')
		line.AppendString(ent.Stack)
	}

	// Coloring
	line.AppendBytes(levelToColorEnd[ent.Level])
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
