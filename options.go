package ulog

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option interface {
	apply(*options)
}

type options struct {
	level                 zap.AtomicLevel
	callerFieldWidth      int
	stripAdditionalFields bool
	renderDummyThread     bool
	replaceNewlines       bool
	newlineReplacement    string
	redirectOutput        io.Writer
	redirectErr           io.Writer
	additionalCores       []zapcore.Core
	baseFields            []zapcore.Field
	errorClassifiers      []func(error) string
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{f: f}
}

// Default: info
func WithLogLevel(lvl string) Option {
	return newFuncOption(func(o *options) {
		level, err := zap.ParseAtomicLevel(lvl)
		if err != nil {
			Errorf("cannot parse level %s, falling back to %s", lvl, zapcore.InfoLevel)
			o.level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
			return
		}
		o.level = level
	})
}

// Default: -1 (no width)
func WithCallerFieldWidth(fieldWidth int) Option {
	return newFuncOption(func(o *options) {
		o.callerFieldWidth = fieldWidth
	})
}

// Default: true
func WithStripAdditionalFields(stripAdditionalFields bool) Option {
	return newFuncOption(func(o *options) {
		o.stripAdditionalFields = stripAdditionalFields
	})
}

// Default: ""
func WithNewlineReplacement(newlineReplacement string) Option {
	return newFuncOption(func(o *options) {
		o.replaceNewlines = true
		o.newlineReplacement = newlineReplacement
	})
}

// Default: os.Stdout
func WithRedirectOutput(w io.Writer) Option {
	return newFuncOption(func(o *options) {
		o.redirectOutput = w
	})
}

// Default: nil
func WithAdditionalCores(cores ...zapcore.Core) Option {
	return newFuncOption(func(o *options) {
		o.additionalCores = cores
	})
}

// WithBaseFields attaches persistent fields to the global logger, so they are
// emitted on every subsequent log entry. They are applied to all cores (the
// console encoder and any additional cores, e.g. an OTEL core), which makes
// this the right place to stamp fleet-wide context such as the build version.
//
// Default: nil
func WithBaseFields(fields ...zapcore.Field) Option {
	return newFuncOption(func(o *options) {
		o.baseFields = fields
	})
}

// WithErrorClassifiers registers extra error classifiers used by ErrorType (and
// therefore by Entry.Err). They are consulted, in order, only after the built-in
// classes find no match; the first non-empty token wins. This lets a service add
// its own low-cardinality classes — e.g. a mongo driver error matched by string
// signature — without editing ulog's core. Each token is a log-dedup bucket, so
// keep them low-cardinality.
//
// Default: nil
func WithErrorClassifiers(fns ...func(error) string) Option {
	return newFuncOption(func(o *options) {
		o.errorClassifiers = fns
	})
}
