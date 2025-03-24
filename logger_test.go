package ulog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	Configure(
		WithLogLevel("trace"),
		WithRedirectOutput(b),
	)
	Info("test")
	lines := lines(t, b)
	require.Len(t, lines, 1)
	requireMsgEquals(t, lines[0], zapcore.InfoLevel, "test")
}
