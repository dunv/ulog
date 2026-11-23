package ulog

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestConsoleKeyValues(t *testing.T) {
	for _, colors := range []bool{false, true} {
		t.Run(map[bool]string{false: "plain", true: "colored"}[colors], func(t *testing.T) {
			enc := newCustomEncoder(options{consoleKeyValues: true, disableColors: !colors, callerFieldWidth: -1})
			fields := []zap.Field{
				zap.String("destination", "inventory"), zap.Int("attempt", 3),
				zap.Error(errors.New("connection refused")),
				zap.String("empty", ""), zap.String("line", "first\nsecond"),
				zap.String("escape", "\x1b[31m"), zap.String("equals", "a=b"),
				zap.Uint64("large", 18446744073709551615),
				zap.Bool("retry", true), zap.String("destination", "other"),
				zap.Any("nested", map[string]any{"items": []any{"a,b", map[string]int{"n": 2}}}),
			}
			buf, err := enc.EncodeEntry(zapcore.Entry{Level: zapcore.WarnLevel, Message: "Failed to connect"}, fields)
			require.NoError(t, err)
			defer buf.Free()
			plain := buf.String()
			if colors {
				require.Contains(t, plain, colorJson+"destination="+colorEnd+"inventory")
				require.Contains(t, plain, colorEnd+" | "+colorJson)
				for _, color := range []string{colorJson, colorEnd, colorWarning} {
					plain = strings.ReplaceAll(plain, color, "")
				}
			}
			require.NotContains(t, plain, "\x1b")
			require.Equal(t, 1, strings.Count(plain, "\n"))
			require.Contains(t, plain, "Failed to connect | destination=inventory attempt=3 error=\"connection refused\" empty=\"\" line=\"first\\nsecond\" escape=\"\\u001b[31m\" equals=\"a=b\" large=18446744073709551615 retry=true destination=other nested={\"items\":[\"a,b\",{\"n\":2}]}\n")
		})
	}
}

func TestConsoleKeyValuesWithFieldsAndAdditionalCore(t *testing.T) {
	var output bytes.Buffer
	core, logs := observer.New(zapcore.DebugLevel)
	Configure(WithRedirectOutput(&output), WithStripAdditionalFields(false), WithConsoleKeyValues(true), WithColors(false), WithAdditionalCores(core), WithBaseFields(zap.String("service", "example")))
	t.Cleanup(func() { Configure() })
	derived := zap.L().With(zap.String("connection", "abc"))
	derived.Warn("failed", zap.Error(errors.New("refused")))
	zap.L().Info("other")
	require.Contains(t, output.String(), "failed | service=example connection=abc error=refused\n")
	require.Contains(t, output.String(), "other | service=example\n")
	require.NotContains(t, strings.Split(output.String(), "other")[1], "connection=")
	require.Equal(t, map[string]any{"service": "example", "connection": "abc", "error": "refused"}, logs.All()[0].ContextMap())
}

func TestConsoleKeyValuesEmptyAndStripped(t *testing.T) {
	for _, stripped := range []bool{false, true} {
		enc := newCustomEncoder(options{consoleKeyValues: true, disableColors: true, stripAdditionalFields: stripped})
		fields := []zap.Field(nil)
		if stripped {
			fields = []zap.Field{zap.String("hidden", "value")}
		}
		buf, err := enc.EncodeEntry(zapcore.Entry{Message: "message"}, fields)
		require.NoError(t, err)
		require.True(t, strings.HasSuffix(buf.String(), " | message\n"), buf.String())
		buf.Free()
	}
}

func TestConsoleKeyValuesNamespaceAndClone(t *testing.T) {
	enc := newCustomEncoder(options{consoleKeyValues: true, disableColors: true})
	zap.String("base", "value").AddTo(enc)
	clone := enc.Clone()
	zap.Namespace("scope").AddTo(clone)
	buf, err := clone.EncodeEntry(zapcore.Entry{Message: "message"}, []zap.Field{zap.String("key", "a,\"b"), zap.Int("n", 2)})
	require.NoError(t, err)
	defer buf.Free()
	require.Contains(t, buf.String(), "base=value scope={\"key\":\"a,\\\"b\",\"n\":2}")
}

func BenchmarkConsoleFields(b *testing.B) {
	cases := []struct {
		name       string
		fields     []zap.Field
		persistent bool
	}{
		{name: "none"},
		{name: "error", fields: []zap.Field{zap.String("destination", "inventory"), zap.Int("attempt", 3), zap.Error(errors.New("connection refused"))}},
		{name: "persistent", persistent: true, fields: []zap.Field{zap.String("service", "example"), zap.String("connection", "abc"), zap.Error(errors.New("connection refused"))}},
		{name: "nested", fields: []zap.Field{zap.Any("payload", map[string]any{"items": []string{"one", "two"}, "count": 2}), zap.String("error", "connection refused")}},
	}
	entry := zapcore.Entry{Time: time.Unix(1700000000, 0), Level: zapcore.WarnLevel, Message: "Failed to connect", Caller: zapcore.NewEntryCaller(0, "/src/service/main.go", 42, true)}
	for _, tc := range cases {
		for _, mode := range []string{"json", "keyvalue", "stripped"} {
			b.Run(tc.name+"/"+mode, func(b *testing.B) {
				enc := newCustomEncoder(options{consoleKeyValues: mode == "keyvalue", stripAdditionalFields: mode == "stripped", callerFieldWidth: -1})
				fields := tc.fields
				if tc.persistent {
					for _, field := range fields {
						field.AddTo(enc)
					}
					fields = nil
				}
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					buf, err := enc.EncodeEntry(entry, fields)
					if err != nil {
						b.Fatal(err)
					}
					buf.Free()
				}
			})
		}
	}
}
