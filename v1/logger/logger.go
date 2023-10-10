package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/clicworth-com/gomicroutils/v1/amqp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

var once sync.Once

var logger *zap.Logger

// Get initializes a zap.Logger instance if it has not been initialized
// already and returns the same instance for subsequent calls.
func Get() *zap.Logger {
	once.Do(func() {
		stdout := zapcore.AddSync(os.Stdout)
		msgq := zapcore.AddSync(amqp.Get())

		level := zap.InfoLevel
		levelEnv := os.Getenv("LOG_LEVEL")
		if levelEnv != "" {
			levelFromEnv, err := zapcore.ParseLevel(levelEnv)
			if err != nil {
				log.Println(
					fmt.Errorf("invalid level, defaulting to INFO: %w", err),
				)
			}

			level = levelFromEnv
		}

		logLevel := zap.NewAtomicLevelAt(level)

		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.TimeKey = "timestamp"
		productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		productionCfg.CallerKey = "caller"
		productionCfg.EncodeCaller = zapcore.FullCallerEncoder

		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
		msgqEncoder := zapcore.NewJSONEncoder(productionCfg)

		sn := os.Getenv("SRV_NAME")
		if sn == "" {
			sn = "dev"
		}
		st := os.Getenv("SRV_TAG")
		if st == "" {
			st = "dev"
		}
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, stdout, logLevel),
			zapcore.NewCore(msgqEncoder, msgq, logLevel).
				With(
					[]zapcore.Field{
						zap.String("srv_name", sn),
						zap.String("srv_tag", st),
					},
				),
		)

		logger = zap.New(core)
	})

	return logger
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, the default logger is returned, unless it is nil
// in which case a disabled logger is returned.
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	} else if l := logger; l != nil {
		return l
	}

	return zap.NewNop()
}

// WithCtx returns a copy of ctx with the Logger attached.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		if lp == l {
			// Do not store same logger.
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, l)
}
