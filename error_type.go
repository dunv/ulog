package ulog

import (
	"context"
	"errors"
	"io"
	"net"
	"sync"
	"syscall"
)

// registered error classifiers (extension point via WithErrorClassifiers),
// published by Configure and read by ErrorType. Guarded independently of
// _globalMu so classification never contends with logger swaps.
var (
	_classifiersMu sync.RWMutex
	_classifiers   []func(error) string
)

func setErrorClassifiers(fns []func(error) string) {
	_classifiersMu.Lock()
	_classifiers = fns
	_classifiersMu.Unlock()
}

// ErrorType maps err to a short, low-cardinality class token ("timeout",
// "conn_reset", "eof", "canceled", …) or "" when err is nil or matches no class.
//
// It exists for log clustering/dedup: a stable message plus a low-cardinality
// class token lets downstream aggregators bucket lines by (call site, class) and
// keep per-class counts, instead of every line being unique because the full
// error string was baked into the message. The Entry builder appends this token
// to the message (see Entry.Err); keep the full error in the "error" field.
//
// Built-in classes are checked first (errors.Is/errors.As based). If none match,
// any classifiers registered via WithErrorClassifiers are consulted in order and
// the first non-empty token wins. Keep the set of tokens small and stable — each
// distinct token is a separate dedup bucket.
func ErrorType(err error) string {
	if err == nil {
		return ""
	}
	if class := builtinErrorType(err); class != "" {
		return class
	}
	_classifiersMu.RLock()
	fns := _classifiers
	_classifiersMu.RUnlock()
	for _, fn := range fns {
		if class := fn(err); class != "" {
			return class
		}
	}
	return ""
}

func builtinErrorType(err error) string {
	switch {
	case errors.Is(err, context.Canceled):
		return "canceled"
	case errors.Is(err, context.DeadlineExceeded):
		return "timeout"
	case errors.Is(err, syscall.ECONNREFUSED):
		return "conn_refused"
	case errors.Is(err, syscall.ECONNRESET):
		return "conn_reset"
	case errors.Is(err, syscall.EPIPE):
		return "broken_pipe"
	case errors.Is(err, syscall.EHOSTUNREACH):
		return "host_unreachable"
	case errors.Is(err, syscall.ENETUNREACH):
		return "net_unreachable"
	case errors.Is(err, net.ErrClosed):
		return "conn_closed"
	case errors.Is(err, io.ErrUnexpectedEOF), errors.Is(err, io.EOF):
		return "eof"
	}
	// Net-level read/write timeouts (e.g. a mongo "i/o timeout") surface as a
	// net.Error with Timeout()==true rather than context.DeadlineExceeded.
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return "timeout"
	}
	return ""
}
