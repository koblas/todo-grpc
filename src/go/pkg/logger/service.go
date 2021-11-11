package logger

import "go.uber.org/zap/zapcore"

type Level zapcore.Level

const (
	LevelDebug Level = Level(zapcore.DebugLevel)
	LevelInfo  Level = Level(zapcore.InfoLevel)
	LevelError Level = Level(zapcore.ErrorLevel)
)

// Logger service interface
type Logger interface {
	With(args ...interface{}) Logger
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	// os.Exit
	Fatal(message string, args ...interface{})
	// golang panic()
	Panic(message string, args ...interface{})

	Latency() func(message string, args ...interface{})
}
