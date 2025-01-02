package otel

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dunv/ulog/v2/otel/otelzap"
	"go.opentelemetry.io/otel/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewOTELCore(otel log.LoggerProvider, lvl string) zapcore.Core {
	level, err := zap.ParseAtomicLevel(lvl)
	if err != nil {
		fmt.Printf("cannot parse level %s, falling back to %s \n", lvl, zapcore.InfoLevel)
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	return &otelCore{
		lvl:      level,
		provider: otel,
	}
}

type otelCore struct {
	lvl      zap.AtomicLevel
	provider log.LoggerProvider
}

func (c *otelCore) Enabled(zapcore.Level) bool {
	return true
}

func (c *otelCore) With(fs []zapcore.Field) zapcore.Core {
	panic("not implemented at the moment")
}

func (c *otelCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *otelCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	rec := log.Record{}
	rec.SetBody(log.StringValue(ent.Message))
	rec.SetTimestamp(ent.Time)
	rec.SetObservedTimestamp(time.Now())
	rec.SetSeverity(otelzap.ConvertLevel(ent.Level))
	rec.SetSeverityText(ent.Level.String())
	rec.AddAttributes(
		log.String("code.function", ent.Caller.Function),
		log.String("code.filepath", ent.Caller.File),
		log.String("code.lineno", strconv.Itoa(ent.Caller.Line)),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c.provider.Logger(ent.LoggerName).Emit(ctx, rec)
	return nil
}

func (c *otelCore) Sync() error {
	return nil
}
