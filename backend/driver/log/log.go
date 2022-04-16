package log

import (
	"context"
	"fmt"
	"time"

	cx "firebase-authentication/driver/context"
	"firebase-authentication/driver/env"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

const format = "2006-01-02T15:04:05.000Z"

func init() {
	level := zap.InfoLevel

	if env.IsDev() {
		level = zap.DebugLevel
	}

	encoder := zap.NewDevelopmentEncoderConfig()

	if env.IsDev() {
		encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoder.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(format))
	})

	zapconf := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoder,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	log, _ = zapconf.Build()
	defer log.Sync()
}

func WithCtx(ctx context.Context) *zap.Logger {
	id := cx.GetReqCtx(ctx)
	return log.Named(id)
}

func Log() *zap.Logger {
	return log
}

func Access(ctx context.Context, path, method string, start time.Time) {
	elapsed := time.Since(start)
	info := fmt.Sprintf("%s %s %s", path, method, elapsed)
	WithCtx(ctx).Info(info)
}

func Elapsed(ctx context.Context, start time.Time, msg string) {
	elapsed := time.Since(start)
	info := fmt.Sprintf("%s elapsed %s", msg, elapsed)
	WithCtx(ctx).Info(info)
}
