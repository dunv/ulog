package v2

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option interface {
	apply(*options)
}

type options struct {
	level                 zap.AtomicLevel
	callerFieldWidth      int
	isDevelopment         bool
	stripAdditionalFields bool
	renderDummyThread     bool
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

func WithCallerFieldWidth(fieldWidth int) Option {
	return newFuncOption(func(o *options) {
		o.callerFieldWidth = fieldWidth
	})
}

func WithIsDevelopment(isDevelopment bool) Option {
	return newFuncOption(func(o *options) {
		o.isDevelopment = isDevelopment
	})
}

func WithStripAdditionalFields(stripAdditionalFields bool) Option {
	return newFuncOption(func(o *options) {
		o.stripAdditionalFields = stripAdditionalFields
	})
}

func WithRenderDummyThread(renderDummyThread bool) Option {
	return newFuncOption(func(o *options) {
		o.renderDummyThread = renderDummyThread
	})
}
