package v2

import (
	"testing"

	"go.uber.org/zap"
)

func TestLogOutput(t *testing.T) {
	zap.S().Debug("logDebug")
	zap.S().Warnw("hello", "field1", "value1")

	zap.L().Warn("hello", zap.Any("field1", "value1"))

	zap.S().Info("logInfo")
	zap.S().Warn("logWarn")
	zap.S().Error("logError")
	zap.S().DPanic("logDPanic")
	childLog := zap.S().With("parent", "context")
	childLog.Infow("testing", "hello", "world")

}
