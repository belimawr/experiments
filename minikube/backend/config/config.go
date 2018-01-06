package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Config struct that holds all configuration parameters
type Config struct {
	ServiceName string `env:"SERVICE" envDefault:"Minikube-Backend"`
	Port        int    `env:"PROT" envDefault:"3000"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"debug"`
	LogOutput   string `env:"LOG_OUTPUT" envDefault:"terminal"`
}

const (
	outputTerminal = "terminal"
)

// ZeroLog returns a zerolog.Logger
func (c Config) ZeroLog() zerolog.Logger {
	var level zerolog.Level

	switch strings.ToLower(c.LogLevel) {
	case "info":
		level = zerolog.InfoLevel
	case "debug":
		level = zerolog.DebugLevel
	case "error":
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}

	logger := zerolog.New(os.Stderr).
		Level(level).
		With().
		Str("service", c.ServiceName).
		Timestamp().
		Logger()

	if c.LogOutput == outputTerminal {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return logger
}
