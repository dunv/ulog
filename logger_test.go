package ulog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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

// WithBaseFields must attach the field to every subsequent entry, regardless of
// the message, so it lands in the structured part of each rendered line.
func TestWithBaseFields(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	Configure(
		WithLogLevel("trace"),
		WithStripAdditionalFields(false),
		WithRedirectOutput(b),
		WithBaseFields(zap.String("version", "1.2.3")),
	)
	Info("first")
	Info("second")
	lines := lines(t, b)
	require.Len(t, lines, 2)
	for _, line := range lines {
		require.Contains(t, string(line), `"version":"1.2.3"`)
	}
}
