package logger

/*
import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/koblas/projectx/server-go/pkg/config"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Entry
}

// NewLogger - Construct a logger, we use a logrus Entry so that it's easy to
// add extra fields as we're passed around
func NewLogrus(config config.Config) Logger {
	log := logrus.New()

	// Only log the warning severity or above.
	level := logrus.InfoLevel

	switch strings.ToUpper(config.LogLevel) {
	case "FATAL":
		level = logrus.FatalLevel
	case "ERROR":
		level = logrus.ErrorLevel
	case "INFO":
		level = logrus.InfoLevel
	case "DEBUG":
		level = logrus.DebugLevel
	}

	log.SetLevel(level)

	if config.SumoEndpoint != "" {
		name, _ := os.Hostname()
		hook := NewSumoLogicHook(config.SumoEndpoint, name, level)
		log.AddHook(hook)
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stdout)
		if config.LogFormat == "json" {
			log.SetFormatter(&logrus.JSONFormatter{})
		} else {
			log.SetFormatter(&logrus.TextFormatter{})
		}
	}

	return &logrusLogger{logrus.NewEntry(log)}
}

var errMissingValue = errors.New("(MISSING)")

func (logger *logrusLogger) With(args ...interface{}) Logger {
	fields := logger.buildFields(args)
	return &logrusLogger{logger.logger.WithFields(fields)}
}

func (logger *logrusLogger) Debug(message string, args ...interface{}) {
	logger.logger.WithFields(logger.buildFields(args)).Debug(message)
}

func (logger *logrusLogger) Info(message string, args ...interface{}) {
	logger.logger.WithFields(logger.buildFields(args)).Info(message)
}
func (logger *logrusLogger) Error(message string, args ...interface{}) {
	logger.logger.WithFields(logger.buildFields(args)).Error(message)
}
func (logger *logrusLogger) Panic(message string, args ...interface{}) {
	logger.logger.WithFields(logger.buildFields(args)).Panic(message)
}
func (logger *logrusLogger) Fatal(message string, args ...interface{}) {
	logger.logger.WithFields(logger.buildFields(args)).Fatal(message)
}

///
func (logger *logrusLogger) buildFields(args []interface{}) logrus.Fields {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			fields[fmt.Sprint(args[i])] = args[i+1]
		} else {
			fields[fmt.Sprint(args[i])] = errMissingValue
		}
	}

	return fields
}

func (logger *logrusLogger) Latency() func(message string, args ...interface{}) {
	start := time.Now()
	return func(message string, args ...interface{}) {
		logger.With(
			"latency", time.Since(start),
		).Info(message, args...)
	}
}

*/
