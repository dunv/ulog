package ulog

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// configureBuf points the global logger at a buffer with structured fields
// rendered (they are stripped by default), so tests can assert the trailing JSON.
func configureBuf(t *testing.T) *bytes.Buffer {
	t.Helper()
	b := bytes.NewBuffer([]byte{})
	Configure(
		WithLogLevel("debug"),
		WithStripAdditionalFields(false),
		WithRedirectOutput(b),
	)
	return b
}

// The reported caller must be the .Msg() call site in this test file, not the
// builder internals in entry.go — proves the AddCallerSkip math.
func TestEntryCallerIsCallSite(t *testing.T) {
	b := configureBuf(t)
	New().Info().Msg("hello")
	out := string(lines(t, b)[0])
	require.Contains(t, out, "ulog/entry_test.go:", "caller should be the test file: %s", out)
	require.NotContains(t, out, "ulog/entry.go:", "caller must not be the builder internals: %s", out)
}

// .Err attaches the full error as a field and appends the classified token to the
// message, keeping the message stable.
func TestEntryErrClassifiesAndFields(t *testing.T) {
	b := configureBuf(t)
	For("serial").Error().Err(context.DeadlineExceeded).F("command", "GET").Msg("command failed")
	out := string(lines(t, b)[0])
	require.Contains(t, out, "[serial] command failed [timeout]")
	require.Contains(t, out, `"command":"GET"`)
	require.Contains(t, out, `"error":"context deadline exceeded"`)
}

// ErrClass overrides the automatic classification.
func TestEntryErrClassOverride(t *testing.T) {
	b := configureBuf(t)
	New().Error().Err(context.Canceled).ErrClass("rate_limited").Msg("upstream rejected")
	out := string(lines(t, b)[0])
	require.Contains(t, out, "upstream rejected [rate_limited]")
	require.NotContains(t, out, "[canceled]")
	require.Contains(t, out, `"error":"context canceled"`)
}

// A message with no error and no class stays exactly as given (no suffix).
func TestEntryNoErrNoClassSuffix(t *testing.T) {
	b := configureBuf(t)
	New().Warn().F("cartID", "140").Msg("cart offline")
	out := string(lines(t, b)[0])
	require.Contains(t, out, " | cart offline")   // New() → no "[label] " prefix
	require.NotContains(t, out, "cart offline [") // no class suffix appended
	require.Contains(t, out, `"cartID":"140"`)
}

// For adds a "[label] " prefix; New adds none.
func TestEntryPrefix(t *testing.T) {
	b := configureBuf(t)
	For("dynamic_config").Info().Msg("watch started")
	out := string(lines(t, b)[0])
	require.Contains(t, out, "[dynamic_config] watch started")
}

// With pre-binds base fields shared across lines.
func TestEntryWithBaseFields(t *testing.T) {
	b := configureBuf(t)
	c := For("sync").With("rule", "products", "batch", 3)
	c.Info().Msg("first")
	c.Error().Err(context.Canceled).Msg("second")
	ls := lines(t, b)
	require.Len(t, ls, 2)
	for _, l := range ls {
		require.Contains(t, string(l), `"rule":"products"`)
		require.Contains(t, string(l), `"batch":3`)
	}
	require.Contains(t, string(ls[1]), "[sync] second [canceled]")
}

// Msgf formats the message (escape hatch); the class suffix still applies.
func TestEntryMsgf(t *testing.T) {
	b := configureBuf(t)
	New().Error().Err(context.DeadlineExceeded).Msgf("attempt %d failed", 4)
	out := string(lines(t, b)[0])
	require.Contains(t, out, "attempt 4 failed [timeout]")
}

// A cheap sanity check that the emitted message segment never carries a raw error
// string — the whole point is that detail lives in the field, not the message.
func TestEntryMessageStaysStable(t *testing.T) {
	b := configureBuf(t)
	longErr := context.DeadlineExceeded
	New().Error().Err(longErr).Msg("load failed")
	out := string(lines(t, b)[0])
	msgSegment := out
	if i := strings.Index(out, " | {"); i >= 0 {
		msgSegment = out[:i] // strip the JSON blob
	}
	require.Contains(t, msgSegment, "load failed [timeout]")
	require.NotContains(t, msgSegment, "context deadline exceeded")
}
