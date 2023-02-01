package v2

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var selectedOptions = options{
	level:                 zap.NewAtomicLevelAt(zapcore.InfoLevel),
	callerFieldWidth:      -1,
	isDevelopment:         false,
	stripAdditionalFields: true,
	renderDummyThread:     false,
}

var skipOneSugaredLogger *zap.SugaredLogger
var skipOneLogger *zap.Logger

func init() {
	reinit()
}

func reinit() {
	encoder := newCustomEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), selectedOptions.level)
	opts := []zap.Option{zap.AddCaller()}
	if selectedOptions.isDevelopment {
		opts = append(opts, zap.Development())
	}
	logger := zap.New(core).WithOptions(opts...)
	zap.ReplaceGlobals(logger)
	skipOneSugaredLogger = zap.S().WithOptions(zap.AddCallerSkip(1))
	skipOneLogger = zap.L().WithOptions(zap.AddCallerSkip(1))
}

func Configure(opts ...Option) {
	for _, opt := range opts {
		opt.apply(&selectedOptions)
	}
	reinit()
}
