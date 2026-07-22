package ulog

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Ctx is a reusable logging context: pre-bound structured fields plus an optional
// "[label] " message prefix, shared by every line logged through it. Build one at
// a call site (cheap) or pass it down to code that must log with the same context.
//
// The convention it enforces is: keep variable data in FIELDS (.F / .Err), not in
// the message. Downstream log shipping (e.g. cart/store log_service) dedups lines by
// exact message text, so a value baked into the message fragments the buckets, while
// a value in a field does not. Terminate with .Msg (constant message — preferred) or
// .Msgf (formatted — escape hatch).
//
// Usage: ulog.For("serial").Error().Err(err).F("command", cmd).Msg("command failed")
type Ctx struct {
	base   []any
	prefix string
}

// New returns a Ctx with no prefix and no base fields.
func New() *Ctx {
	return &Ctx{}
}

// For returns a Ctx that renders a "[label] " prefix in front of every message.
// Use New for no prefix.
func For(label string) *Ctx {
	return &Ctx{prefix: "[" + label + "] "}
}

// With returns a copy of c with additional pre-bound key/value fields. Fields are
// key, value, key, value, … just like the ...w logging functions.
func (c *Ctx) With(kv ...any) *Ctx {
	if len(kv) == 0 {
		return c
	}
	base := make([]any, 0, len(c.base)+len(kv))
	base = append(base, c.base...)
	base = append(base, kv...)
	return &Ctx{base: base, prefix: c.prefix}
}

// Per-level entry constructors. They return an *Entry to chain .Err/.ErrClass/.F/
// .Fields and terminate with .Msg/.Msgf. Using these instead of a level argument
// keeps zapcore out of call sites.
func (c *Ctx) Debug() *Entry { return &Entry{c: c, lvl: zapcore.DebugLevel} }
func (c *Ctx) Info() *Entry  { return &Entry{c: c, lvl: zapcore.InfoLevel} }
func (c *Ctx) Warn() *Entry  { return &Entry{c: c, lvl: zapcore.WarnLevel} }
func (c *Ctx) Error() *Entry { return &Entry{c: c, lvl: zapcore.ErrorLevel} }
func (c *Ctx) Fatal() *Entry { return &Entry{c: c, lvl: zapcore.FatalLevel} }

// Entry is a single in-progress log line. Not safe for concurrent use — build and
// terminate it inline.
type Entry struct {
	c        *Ctx
	lvl      zapcore.Level
	err      error
	class    string
	classSet bool
	fields   []any
}

// Err attaches err as a dedicated "error" field and, unless an explicit class was
// set with ErrClass, appends a low-cardinality class token (e.g. " [timeout]") to
// the message. The full error rides in the field so the message stays stable for
// dedup. A nil err is a no-op. See ErrorType.
func (e *Entry) Err(err error) *Entry {
	e.err = err
	return e
}

// ErrClass sets the message's class token explicitly, overriding ErrorType. Use it
// for errors errors.Is/errors.As can't reach (e.g. a driver error matched by string)
// when registering a WithErrorClassifiers classifier isn't warranted. Keep the token
// low-cardinality — each distinct token is a separate dedup bucket.
func (e *Entry) ErrClass(token string) *Entry {
	e.class = token
	e.classSet = true
	return e
}

// F attaches a single structured key/value field. Put variable data here, not in
// the message.
func (e *Entry) F(key string, val any) *Entry {
	e.fields = append(e.fields, key, val)
	return e
}

// Fields attaches key, value, key, value, … structured fields in bulk.
func (e *Entry) Fields(kv ...any) *Entry {
	e.fields = append(e.fields, kv...)
	return e
}

// Msg emits the entry with a constant message. Preferred: keeps every emission of
// this call site on one message string so downstream dedup can merge them.
func (e *Entry) Msg(msg string) {
	e.emit(msg)
}

// Msgf emits the entry with a formatted message. Escape hatch — only for
// low-cardinality formatting or non-forwarded (INFO/DEBUG) lines; a high-cardinality
// value baked in here defeats exact-match dedup. Prefer Msg + F.
func (e *Entry) Msgf(format string, args ...any) {
	e.emit(fmt.Sprintf(format, args...))
}

func (e *Entry) emit(msg string) {
	// Collect fields: context base, then explicit, then the error.
	n := len(e.fields)
	if e.c != nil {
		n += len(e.c.base)
	}
	if e.err != nil {
		n += 2
	}
	fields := make([]any, 0, n)
	if e.c != nil {
		fields = append(fields, e.c.base...)
	}
	fields = append(fields, e.fields...)
	if e.err != nil {
		fields = append(fields, "error", e.err.Error())
	}

	// Determine the class token: explicit ErrClass wins, else classify the error.
	class := e.class
	if !e.classSet && e.err != nil {
		class = ErrorType(e.err)
	}

	// Assemble "[prefix] msg [class]".
	var b strings.Builder
	if e.c != nil {
		b.WriteString(e.c.prefix)
	}
	b.WriteString(msg)
	if class != "" {
		b.WriteString(" [")
		b.WriteString(class)
		b.WriteByte(']')
	}

	// AddCallerSkip(1) covers this emit + the Msg/Msgf frame on top of SkipOneS's
	// baseline, so the log carries the real call site rather than entry.go.
	SkipOneS().WithOptions(zap.AddCallerSkip(1)).Logw(e.lvl, b.String(), fields...)
}
