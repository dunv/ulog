package ulog

import (
	"fmt"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	colorJson    string = "\033[38:5:237m"
	colorDebug   string = "\033[90m"
	colorWarning string = "\033[93m"
	colorError   string = "\033[91m"
	colorFatal   string = "\033[95m"
	colorEnd     string = "\033[0m"
	timeFormat          = time.RFC3339
)

var (
	levelToColorStart = map[zapcore.Level][]byte{
		zapcore.DebugLevel:  []byte(colorDebug),
		zapcore.InfoLevel:   nil,
		zapcore.WarnLevel:   []byte(colorWarning),
		zapcore.ErrorLevel:  []byte(colorError),
		zapcore.DPanicLevel: []byte(colorFatal),
		zapcore.PanicLevel:  []byte(colorFatal),
		zapcore.FatalLevel:  []byte(colorFatal),
	}
	levelToColorEnd = map[zapcore.Level][]byte{
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

var _pool = buffer.NewPool()

func init() {
	for level := range levelToColorStart {
		lvlStr := level.CapitalString()
		levelString[level] = fmt.Sprintf("%-6s", lvlStr)
		if levelFieldWidth < len(lvlStr) {
			levelFieldWidth = len(lvlStr)
		}
	}
}

// All other methods for adding fields are exposed through zapcore.Encoder interface,
type customEncoder struct {
	zapcore.Encoder
	cfg  zapcore.EncoderConfig
	opts options
}

func newCustomEncoder(opts options) zapcore.Encoder {
	// use empty keys to avoid rendering them in the output
	jsonEncoderCfg := zapcore.EncoderConfig{
		TimeKey:          "",
		LevelKey:         "",
		NameKey:          "",
		CallerKey:        "",
		FunctionKey:      "",
		MessageKey:       "",
		StacktraceKey:    "",
		SkipLineEnding:   true,
		ConsoleSeparator: " | ",
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout(timeFormat),
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
	}

	return &customEncoder{
		zapcore.NewJSONEncoder(jsonEncoderCfg),
		jsonEncoderCfg,
		opts,
	}
}

func (c customEncoder) Clone() zapcore.Encoder {
	return &customEncoder{
		c.Encoder.Clone(),
		c.cfg,
		c.opts,
	}
}

func (c customEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	line := _pool.Get()

	// Coloring
	if !c.opts.disableColors {
		line.AppendBytes(levelToColorStart[ent.Level])
	}
	line.AppendString(ent.Time.Format(timeFormat))
	line.AppendString(c.cfg.ConsoleSeparator)
	c.appendPaddedLevel(ent.Level, line)
	line.AppendString(c.cfg.ConsoleSeparator)
	c.appendCaller(ent.Caller, line)
	line.AppendString(c.cfg.ConsoleSeparator)

	if c.opts.replaceNewlines {
		line.AppendString(strings.ReplaceAll(ent.Message, "\n", c.opts.newlineReplacement))
	} else {
		line.AppendString(ent.Message)
	}

	// cannot check for field-presence here: if log.With is used, the fields will be added
	// BEFORE EncodeEntry is called with zero fields
	if !c.opts.stripAdditionalFields {
		// "Abuse" zapcore's jsonEncoder to render fields
		buf, _ := c.Encoder.EncodeEntry(ent, fields)
		if buf.Len() > 2 { // len() == 2 means only "{}" was written, no fields
			if c.opts.consoleKeyValues {
				if !c.opts.disableColors {
					line.AppendString(colorEnd)
				}
				line.AppendString(c.cfg.ConsoleSeparator)
				c.appendKeyValues(line, buf.Bytes())
			} else {
				if !c.opts.disableColors {
					line.AppendString(colorJson)
				}
				line.AppendString(c.cfg.ConsoleSeparator)
				_, _ = line.Write(buf.Bytes())
			}
			if !c.opts.disableColors {
				line.AppendString(colorEnd)
			}
		} else if !c.opts.disableColors {
			line.AppendBytes(levelToColorEnd[ent.Level])
		}
		buf.Free()
		line.AppendString(zapcore.DefaultLineEnding)

		// return here, so we do not have to remember if json was added or not
		return line, nil
	}

	// Coloring
	if !c.opts.disableColors {
		line.AppendBytes(levelToColorEnd[ent.Level])
	}
	line.AppendString(zapcore.DefaultLineEnding)

	return line, nil
}

// appendKeyValues scans the JSON encoder's trusted output without decoding it
// into maps or interfaces. This preserves order, duplicate keys, numeric precision
// and Zap's field/marshaler semantics while avoiding per-field allocations.
func (c customEncoder) appendKeyValues(line *buffer.Buffer, object []byte) {
	for i := 1; i < len(object)-1; {
		keyEnd := jsonStringEnd(object, i)
		key := consoleValue(object[i:keyEnd])
		i = keyEnd + 1 // skip colon
		start := i
		depth := 0
		for i < len(object)-1 {
			switch object[i] {
			case '"':
				i = jsonStringEnd(object, i)
				continue
			case '{', '[':
				depth++
			case '}', ']':
				depth--
			case ',':
				if depth == 0 {
					goto valueEnd
				}
			}
			i++
		}
	valueEnd:
		if !c.opts.disableColors {
			line.AppendString(colorJson)
		}
		line.AppendBytes(key)
		line.AppendByte('=')
		if !c.opts.disableColors {
			line.AppendString(colorEnd)
		}
		line.AppendBytes(consoleValue(object[start:i]))
		i++ // skip comma or closing brace
		if i < len(object)-1 {
			line.AppendByte(' ')
		}
	}
}

// jsonStringEnd returns the position after a quoted string in trusted JSON.
func jsonStringEnd(s []byte, start int) int {
	for i := start + 1; i < len(s); i++ {
		switch s[i] {
		case '\\':
			i++
		case '"':
			return i + 1
		}
	}
	return len(s)
}

// Keep JSON escaping for special strings, including control characters and ANSI
// escapes. Ordinary strings can be displayed without quotes or allocations.
func consoleValue(raw []byte) []byte {
	if len(raw) < 2 || raw[0] != '"' {
		return raw
	}
	value := raw[1 : len(raw)-1]
	if len(value) == 0 {
		return raw
	}
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case '\\', '"', '=', '{', '}', '[', ']':
			return raw
		}
		if value[i] <= ' ' {
			return raw
		}
		if value[i] >= utf8.RuneSelf {
			r, size := utf8.DecodeRune(value[i:])
			if unicode.IsSpace(r) || unicode.IsControl(r) {
				return raw
			}
			i += size - 1
		}
	}
	return value
}

func (c *customEncoder) appendPaddedLevel(l zapcore.Level, enc *buffer.Buffer) {
	if s, ok := levelString[l]; ok {
		enc.AppendString(s)
	}
}

// copied from zapcore with constant width
func (c *customEncoder) appendCaller(caller zapcore.EntryCaller, enc *buffer.Buffer) {
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
	if c.opts.callerFieldWidth != -1 {
		if len(str) > c.opts.callerFieldWidth {
			str = str[len(str)-c.opts.callerFieldWidth:]
		} else {
			str = str + strings.Repeat(" ", c.opts.callerFieldWidth-len(str))
		}
	}
	enc.AppendString(str)
}
