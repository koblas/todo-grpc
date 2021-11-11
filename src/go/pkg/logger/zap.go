package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

// NewLogger - Construct a logger, we use a zap Entry so that it's easy to
// add extra fields as we're passed around
func NewZap(logLevel Level) Logger {

	// Only log the warning severity or above.
	level := zapcore.Level(logLevel)

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.DisableCaller = true

	logger, _ := cfg.Build()

	sugar := logger.Sugar()

	return &zapLogger{sugar}
}

func (logger *zapLogger) With(args ...interface{}) Logger {
	return &zapLogger{logger.logger.With(args...)}
}

func (logger *zapLogger) Debug(message string, args ...interface{}) {
	logger.logger.Debugw(message, args...)
}

func (logger *zapLogger) Info(message string, args ...interface{}) {
	logger.logger.Infow(message, args...)
}
func (logger *zapLogger) Error(message string, args ...interface{}) {
	logger.logger.Errorw(message, args...)
}
func (logger *zapLogger) Panic(message string, args ...interface{}) {
	logger.logger.Panicw(message, args...)
}
func (logger *zapLogger) Fatal(message string, args ...interface{}) {
	logger.logger.Fatalw(message, args...)
}

func (logger *zapLogger) Latency() func(message string, args ...interface{}) {
	start := time.Now()
	return func(message string, args ...interface{}) {
		logger.With(
			"latency", time.Since(start),
		).Info(message, args...)
	}
}
