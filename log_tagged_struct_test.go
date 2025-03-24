package ulog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestLogJSONStruct(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	Configure(
		WithLogLevel("trace"),
		WithRedirectOutput(b),
	)

	type myStruct struct {
		TestOne string `json:"testOne"`
		TestTwo int    `json:"testTwo"`
	}

	s := myStruct{
		TestOne: "testOne",
		TestTwo: 2,
	}

	LogJSONStruct(s, "funPrefix ")
	lines := lines(t, b)
	require.Len(t, lines, 2)
	requireMsgEquals(t, lines[0], zapcore.InfoLevel, "funPrefix testOne = testOne")
	requireMsgEquals(t, lines[1], zapcore.InfoLevel, "funPrefix testTwo = 2")
}

func TestLogEnvStruct(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	Configure(
		WithLogLevel("trace"),
		WithRedirectOutput(b),
	)

	type myStruct struct {
		Hello  string `env:"hello"`
		World  string `env:"world" warnIf:"world"`
		Masked string `env:"masked" mask:"true"`
	}

	t.Run("with warning", func(t *testing.T) {
		s := myStruct{
			Hello:  "hello",
			World:  "world",
			Masked: "masked",
		}
		LogEnvStruct(s, "\t")
		lines := lines(t, b)
		require.Len(t, lines, 6)
		requireMsgEquals(t, lines[0], zapcore.WarnLevel, "WARNINGS ..............................................................")
		requireMsgEquals(t, lines[1], zapcore.WarnLevel, "\tworld has warning value")
		requireMsgEquals(t, lines[2], zapcore.InfoLevel, "Config ................................................................")
		requireMsgEquals(t, lines[3], zapcore.InfoLevel, "\thello = hello")
		requireMsgEquals(t, lines[4], zapcore.InfoLevel, "\tworld = world")
		requireMsgEquals(t, lines[5], zapcore.InfoLevel, "\tmasked = ***")
		b.Reset()
	})
	t.Run("without warning", func(t *testing.T) {
		s := myStruct{
			Hello:  "hello",
			World:  "world2",
			Masked: "masked",
		}
		LogEnvStruct(s, "\t")
		lines := lines(t, b)
		require.Len(t, lines, 4)
		requireMsgEquals(t, lines[0], zapcore.InfoLevel, "Config ................................................................")
		requireMsgEquals(t, lines[1], zapcore.InfoLevel, "\thello = hello")
		requireMsgEquals(t, lines[2], zapcore.InfoLevel, "\tworld = world2")
		requireMsgEquals(t, lines[3], zapcore.InfoLevel, "\tmasked = ***")
		b.Reset()
	})

}
