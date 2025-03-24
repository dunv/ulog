package ulog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestLogStruct(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	Configure(
		WithLogLevel("trace"),
		WithRedirectOutput(b),
	)

	type myStruct struct {
		Hello string
		World string
	}

	t.Run("to trace", func(t *testing.T) {
		s := myStruct{
			Hello: "hello",
			World: "world",
		}
		TraceStruct(s, "\t")
		lines := lines(t, b)
		require.Len(t, lines, 2)
		requireMsgEquals(t, lines[0], zapcore.TraceLevel, "\tHello = hello")
		requireMsgEquals(t, lines[1], zapcore.TraceLevel, "\tWorld = world")
		b.Reset()
	})
	t.Run("to debug", func(t *testing.T) {
		s := myStruct{
			Hello: "hello",
			World: "world",
		}
		DebugStruct(s, "\t")
		lines := lines(t, b)
		require.Len(t, lines, 2)
		requireMsgEquals(t, lines[0], zapcore.DebugLevel, "\tHello = hello")
		requireMsgEquals(t, lines[1], zapcore.DebugLevel, "\tWorld = world")
		b.Reset()
	})
	t.Run("to warn", func(t *testing.T) {
		s := myStruct{
			Hello: "hello",
			World: "world",
		}
		logStruct(zapcore.WarnLevel, s, "\t")
		lines := lines(t, b)
		require.Len(t, lines, 2)
		requireMsgEquals(t, lines[0], zapcore.WarnLevel, "\tHello = hello")
		requireMsgEquals(t, lines[1], zapcore.WarnLevel, "\tWorld = world")
		b.Reset()
	})
	t.Run("to error", func(t *testing.T) {
		s := myStruct{
			Hello: "hello",
			World: "world",
		}
		ErrorStruct(s, "\t")
		lines := lines(t, b)
		require.Len(t, lines, 2)
		requireMsgEquals(t, lines[0], zapcore.ErrorLevel, "\tHello = hello")
		requireMsgEquals(t, lines[1], zapcore.ErrorLevel, "\tWorld = world")
		b.Reset()
	})

}
