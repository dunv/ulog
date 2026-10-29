package ulog

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestCoreWrapperAppliesToCombinedCoreAndCachedLoggers(t *testing.T) {
	var out bytes.Buffer
	additional, observed := observer.New(zapcore.DebugLevel)
	wrapped := 0
	var hooked []string
	Configure(
		WithLogLevel("info"),
		WithRedirectOutput(&out),
		WithStripAdditionalFields(false),
		WithAdditionalCores(additional),
		WithBaseFields(zap.String("version", "test-version")),
		WithCoreWrapper(func(core zapcore.Core) zapcore.Core {
			wrapped++
			return zapcore.RegisterHooks(core.With([]zap.Field{zap.String("wrapped", "yes")}),
				func(entry zapcore.Entry) error {
					hooked = append(hooked, entry.Message)
					return nil
				})
		}),
	)
	t.Cleanup(func() { Configure() })

	Debug("additional-only")
	Info("sugar")
	For("component").Info().F("detail", "kept").Msg("fluent")
	zap.L().With(zap.String("child", "bound")).Info("derived")

	require.Equal(t, 1, wrapped, "wrap once at configuration, not once per logger or entry")
	require.Equal(t, []string{"additional-only", "sugar", "[component] fluent", "derived"}, hooked)
	require.NotContains(t, out.String(), "additional-only", "wrapping must retain each core's level")
	require.Len(t, lines(t, &out), 3)
	require.Equal(t, 4, observed.Len())
	for _, entry := range observed.All() {
		require.Equal(t, "test-version", entry.ContextMap()["version"])
		require.Equal(t, "yes", entry.ContextMap()["wrapped"])
	}
	require.Equal(t, "kept", observed.FilterMessage("[component] fluent").All()[0].ContextMap()["detail"])
	require.Equal(t, "bound", observed.FilterMessage("derived").All()[0].ContextMap()["child"])
}

func TestCoreWrapperRunsAfterBaseFields(t *testing.T) {
	var out bytes.Buffer
	baseMarshaled := 0
	Configure(
		WithRedirectOutput(&out),
		WithStripAdditionalFields(false),
		WithBaseFields(zap.Object("base", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			baseMarshaled++
			enc.AddString("identity", "trusted")
			return nil
		}))),
		WithCoreWrapper(func(core zapcore.Core) zapcore.Core {
			require.Equal(t, 1, baseMarshaled, "trusted base fields must already be attached before wrapping")
			return core
		}),
	)
	t.Cleanup(func() { Configure() })

	Info("base-proof")
	require.Equal(t, 1, baseMarshaled)
	require.Contains(t, out.String(), `"base":{"identity":"trusted"}`)
}

func TestCoreWrapperDoesNotAddConsoleEncoding(t *testing.T) {
	var out bytes.Buffer
	Configure(
		WithRedirectOutput(&out),
		WithStripAdditionalFields(false),
		WithCoreWrapper(func(core zapcore.Core) zapcore.Core { return core }),
	)
	t.Cleanup(func() { Configure() })

	for _, emit := range []struct {
		name string
		log  func(zapcore.ObjectMarshaler)
	}{
		{"zap", func(payload zapcore.ObjectMarshaler) { zap.L().Info("zap", zap.Object("payload", payload)) }},
		{"sugar", func(payload zapcore.ObjectMarshaler) { NewDefaultLogger().Infow("sugar", "payload", payload) }},
		{"fluent", func(payload zapcore.ObjectMarshaler) { New().Info().F("payload", payload).Msg("fluent") }},
	} {
		t.Run(emit.name, func(t *testing.T) {
			marshaled := 0
			emit.log(zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				marshaled++
				enc.AddString("value", "once")
				return nil
			}))
			require.Equal(t, 1, marshaled, "one console sink must encode the structured payload exactly once")
		})
	}
	require.Equal(t, 3, strings.Count(out.String(), `"payload":{"value":"once"}`))
	require.Len(t, lines(t, &out), 3)
}

func TestCoreWrapperLastOptionWinsAndConfigureResets(t *testing.T) {
	var out bytes.Buffer
	first, second := 0, 0
	firstWrapper := func(core zapcore.Core) zapcore.Core {
		first++
		return core
	}
	Configure(
		WithRedirectOutput(&out),
		WithStripAdditionalFields(false),
		WithCoreWrapper(firstWrapper),
		WithCoreWrapper(func(core zapcore.Core) zapcore.Core {
			second++
			return core.With([]zap.Field{zap.String("wrapper", "second")})
		}),
	)
	t.Cleanup(func() { Configure() })
	Info("wrapped")

	Configure(WithRedirectOutput(&out), WithStripAdditionalFields(false),
		WithCoreWrapper(firstWrapper), WithCoreWrapper(nil))
	Info("nil-wrapper")
	Configure(WithRedirectOutput(&out), WithStripAdditionalFields(false))
	Info("reset")

	require.Zero(t, first)
	require.Equal(t, 1, second)
	require.Equal(t, 1, strings.Count(out.String(), `"wrapper":"second"`))
	require.Len(t, lines(t, &out), 3)
}
