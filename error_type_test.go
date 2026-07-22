package ulog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

// fakeTimeoutErr is a net.Error reporting a timeout, like a mongo "i/o timeout".
type fakeTimeoutErr struct{}

func (fakeTimeoutErr) Error() string   { return "i/o timeout" }
func (fakeTimeoutErr) Timeout() bool   { return true }
func (fakeTimeoutErr) Temporary() bool { return false }

func TestErrorTypeBuiltins(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, ""},
		{"unknown", errors.New("something odd"), ""},
		{"canceled", context.Canceled, "canceled"},
		{"deadline", context.DeadlineExceeded, "timeout"},
		{"conn_refused", syscall.ECONNREFUSED, "conn_refused"},
		{"conn_reset", syscall.ECONNRESET, "conn_reset"},
		{"broken_pipe", syscall.EPIPE, "broken_pipe"},
		{"host_unreachable", syscall.EHOSTUNREACH, "host_unreachable"},
		{"net_unreachable", syscall.ENETUNREACH, "net_unreachable"},
		{"conn_closed", net.ErrClosed, "conn_closed"},
		{"eof", io.EOF, "eof"},
		{"unexpected_eof", io.ErrUnexpectedEOF, "eof"},
		{"net_timeout", fakeTimeoutErr{}, "timeout"},
		{"wrapped_canceled", fmt.Errorf("collect status: %w", context.Canceled), "canceled"},
	}
	// Ensure no classifiers leak in from another test.
	Configure()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, ErrorType(tc.err))
		})
	}
}

// WithErrorClassifiers extends ErrorType for errors the built-ins don't match;
// built-in classes still take precedence.
func TestErrorTypeClassifiersExtension(t *testing.T) {
	sentinel := errors.New("New cluster time is too far from this node's wall clock")
	Configure(WithErrorClassifiers(func(err error) string {
		if err != nil && err.Error() == sentinel.Error() {
			return "cluster_time_skew"
		}
		return "shadow" // would match everything if consulted
	}))
	defer Configure() // reset registered classifiers for later tests

	require.Equal(t, "cluster_time_skew", ErrorType(sentinel))
	// Built-in match wins over the (everything-matching) classifier.
	require.Equal(t, "canceled", ErrorType(context.Canceled))
	// nil is always "".
	require.Equal(t, "", ErrorType(nil))
}

// Reconfiguring without the option clears previously registered classifiers.
func TestErrorTypeClassifiersReset(t *testing.T) {
	Configure(WithErrorClassifiers(func(error) string { return "x" }))
	require.Equal(t, "x", ErrorType(errors.New("anything")))
	Configure()
	require.Equal(t, "", ErrorType(errors.New("anything")))
}
