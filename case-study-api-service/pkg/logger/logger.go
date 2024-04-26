package logger

import (
	"case-study-api-service/pkg/config"
	"context"
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	environmentTypeProduction = "production"
	environmentTypeConsole    = "console"
)

type ctxKey struct{}

var logger = defaultSugarLogger()

type LoggerOption func() zapcore.Core

var once sync.Once

func GetLogger() *zap.Logger {
	return logger
}

func InitLogger(loggerName string, options ...LoggerOption) {
	once.Do(func() {
		var logOptions []zapcore.Core

		for _, opt := range options {
			logOptions = append(logOptions, opt())
		}

		core := zapcore.NewTee(logOptions...)

		l := zap.New(core, zap.AddCaller()).Named(loggerName)

		// set logger
		logger = l
	})
}

func WithIO(w io.Writer, logLevel int, environmentType, logEncoding string) LoggerOption {
	return func() zapcore.Core {
		var (
			encoderConfig zapcore.EncoderConfig
			encoder       zapcore.Encoder
		)

		if environmentType == environmentTypeProduction {
			encoderConfig = zap.NewProductionEncoderConfig()
		} else {
			encoderConfig = zap.NewDevelopmentEncoderConfig()
		}

		if logEncoding == environmentTypeConsole {
			encoder = zapcore.NewConsoleEncoder(encoderConfig)
		} else {
			encoder = zapcore.NewJSONEncoder(encoderConfig)
		}

		stdoutLogger := zapcore.NewCore(
			encoder,
			zapcore.AddSync(w),
			zap.NewAtomicLevelAt(zapcore.Level(logLevel)),
		)

		return stdoutLogger
	}
}

func defaultSugarLogger() *zap.Logger {
	return zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentConfig().EncoderConfig),
		zapcore.AddSync(os.Stderr),
		zap.DebugLevel,
	))
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, the default logger is returned, unless it is nil
// in which case a disabled logger is returned.
func FromCtx(ctx context.Context) *zap.Logger {
	if ctx != nil {
		if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
			return l
		} else if l := logger; l != nil {
			return l
		}
	}

	return defaultSugarLogger()
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

func RegisterLogger(c config.LoggerConfig) *zap.Logger {
	var loggerOpts []LoggerOption

	loggerOpts = append(loggerOpts, WithIO(os.Stdout, c.LogLevel, c.EnvironmentType, c.LogEncoding))

	InitLogger(c.AppName, loggerOpts...)

	return GetLogger()
}
