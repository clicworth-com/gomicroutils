package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/clicworth-com/gomicroutils/v3/amqp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

var once sync.Once

type CBLogger struct {
	logger *zap.Logger
}

var cblog *CBLogger

const FUNCTION_SKIP_LEVEL = 1

// Get initializes a zap.Logger instance if it has not been initialized
// already and returns the same instance for subsequent calls.
func Get() *CBLogger {
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
		productionCfg.EncodeCaller = zapcore.ShortCallerEncoder

		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		developmentCfg.CallerKey = "caller"
		developmentCfg.EncodeCaller = zapcore.ShortCallerEncoder

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

		logger := zap.New(core)
		cblog = &CBLogger{
			logger: logger,
		}
	})

	return cblog
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, the default logger is returned, unless it is nil
// in which case a disabled logger is returned.
func FromCtx(ctx context.Context) *CBLogger {
	if l, ok := ctx.Value(ctxKey{}).(*CBLogger); ok {
		return l
	} else if l := cblog; l != nil {
		return l
	}

	return &CBLogger{zap.NewNop()}
}

// WithCtx returns a copy of ctx with the Logger attached.
func WithCtx(ctx context.Context, l *CBLogger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*CBLogger); ok {
		if lp == l {
			// Do not store same logger.
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, l)
}

// WithCtx returns a copy of ctx with the Logger attached.
func (l *CBLogger) With(f zapcore.Field) *CBLogger {
	nl := l.logger.With(f)
	return &CBLogger{nl}
}

func (c *CBLogger) Info(msg string, fields ...zap.Field) {

	pc, file, line, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		filetok := strings.Split(file, "/")
		filename := filetok[len(filetok)-1]
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		funcname := functok[len(functok)-1]
		caller := filename + "->" + funcname + ":" + fmt.Sprint(line)
		fields = append(fields, zap.String("caller", caller))
	}
	c.logger.Info(msg, fields...)
}

func (c *CBLogger) Debug(msg string, fields ...zap.Field) {

	pc, file, line, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		filetok := strings.Split(file, "/")
		filename := filetok[len(filetok)-1]
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		funcname := functok[len(functok)-1]
		caller := filename + "->" + funcname + ":" + fmt.Sprint(line)
		fields = append(fields, zap.String("caller", caller))
	}
	c.logger.Debug(msg, fields...)

}

func (c *CBLogger) Warn(msg string, fields ...zap.Field) {

	pc, file, line, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		filetok := strings.Split(file, "/")
		filename := filetok[len(filetok)-1]
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		funcname := functok[len(functok)-1]
		caller := filename + "->" + funcname + ":" + fmt.Sprint(line)
		fields = append(fields, zap.String("caller", caller))
	}
	c.logger.Warn(msg, fields...)

}

func (c *CBLogger) Error(msg string, fields ...zap.Field) {
	pc, file, line, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		filetok := strings.Split(file, "/")
		filename := filetok[len(filetok)-1]
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		funcname := functok[len(functok)-1]
		caller := filename + "->" + funcname + ":" + fmt.Sprint(line)
		fields = append(fields, zap.String("caller", caller))
	}

	c.logger.Error(msg, fields...)

}

func (c *CBLogger) Panic(msg string, fields ...zap.Field) {
	pc, file, line, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		filetok := strings.Split(file, "/")
		filename := filetok[len(filetok)-1]
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		funcname := functok[len(functok)-1]
		caller := filename + "->" + funcname + ":" + fmt.Sprint(line)
		fields = append(fields, zap.String("caller", caller))
	}

	c.logger.Panic(msg, fields...)

}

func (c *CBLogger) Fatal(msg string, fields ...zap.Field) {
	pc, file, line, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		filetok := strings.Split(file, "/")
		filename := filetok[len(filetok)-1]
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		funcname := functok[len(functok)-1]
		caller := filename + "->" + funcname + ":" + fmt.Sprint(line)
		fields = append(fields, zap.String("caller", caller))
	}

	c.logger.Fatal(msg, fields...)

}
