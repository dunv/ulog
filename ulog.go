package ulog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	Configure()
}

func Configure(o ...Option) {
	opts := options{
		level:                 zap.NewAtomicLevelAt(zapcore.InfoLevel),
		callerFieldWidth:      -1,
		stripAdditionalFields: true,
		renderDummyThread:     false,
		replaceNewlines:       false,
		newlineReplacement:    "",
		redirectOutput:        os.Stdout,
		redirectErr:           os.Stderr,
		additionalCores:       nil,
	}

	for _, opt := range o {
		opt.apply(&opts)
	}
	encoder := newCustomEncoder(opts)
	sink := zapcore.AddSync(opts.redirectOutput)
	core := zapcore.NewCore(encoder, sink, opts.level)

	zapOpts := []zap.Option{
		zap.ErrorOutput(zapcore.AddSync(opts.redirectErr)),
		zap.AddCaller(),
	}

	if len(opts.additionalCores) > 0 {
		cores := append([]zapcore.Core{core}, opts.additionalCores...)
		core = zapcore.NewTee(cores...)
	}

	logger := zap.New(core, zapOpts...)
	if len(opts.baseFields) > 0 {
		logger = logger.With(opts.baseFields...)
	}
	zap.ReplaceGlobals(logger)

	_globalMu.Lock()
	_globalL = zap.L().WithOptions(zap.AddCallerSkip(1))
	_globalS = _globalL.Sugar()
	_globalMu.Unlock()
}
