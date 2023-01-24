package v2

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var skipOneLogger *zap.SugaredLogger

func init() {
	reinit(zap.NewAtomicLevelAt(zapcore.InfoLevel), false)
}

func reinit(lvl zap.AtomicLevel, isDevelopment bool) {
	var logger *zap.Logger
	if isDevelopment {
		encoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
		core := zapcore.NewTee(zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lvl))
		logger = zap.New(core).WithOptions(zap.AddCaller(), zap.Development())
	} else {
		encoder := zapcore.NewJSONEncoder(jsonEncoderConfig)
		core := zapcore.NewTee(zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lvl))
		logger = zap.New(core).WithOptions(zap.AddCaller())
	}

	zap.ReplaceGlobals(logger)
	skipOneLogger = zap.S().WithOptions(zap.AddCallerSkip(1))
}

func Configure(logLevel string, isDevelopment bool) error {
	if strings.ToUpper(logLevel) == "TRACE" {
		zap.S().Warnf("Level TRACE is deprecated, using DEBUG instead")
		logLevel = "debug"
	}

	lvl, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		return err
	}
	reinit(lvl, isDevelopment)
	return nil
}
