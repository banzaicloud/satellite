package config

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var logger *logrus.Logger
var loggerOnce sync.Once

// Logger is a configured Logrus logger
func Logger() *logrus.Logger {
	loggerOnce.Do(func() { logger = newLogger() })

	return logger
}

// Config holds information necessary for customizing the logger.
type Config struct {
	Level  string
	Format string
}

func newLogger() *logrus.Logger {
	logger := NewLogger(Config{
		Level:  "debug",
		Format: "",
	})

	return logger
}

// NewLogger creates a new logrus logger instance.
func NewLogger(config Config) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}

	logger.Level = level

	switch config.Format {
	case "json":
		logger.Formatter = new(logrus.JSONFormatter)

	default:
		textFormatter := new(logrus.TextFormatter)
		textFormatter.FullTimestamp = true

		logger.Formatter = textFormatter
	}

	return logger
}
