package slogger

import (
	"errors"
	"log/slog"
	"os"
	"strings"
)

var (
	ErrInvalidLogLevel = errors.New("log level not recognized, falling back to info")
	ErrInvalidLogFormat = errors.New("log format not recognized, falling back to text")
)

// NewLogger returns a logger with the specified log level and format
func NewLogger(level, format string) (logger *slog.Logger, err error) {
	var slogHandlerOptions slog.HandlerOptions
	var slogHandler slog.Handler

	switch strings.ToLower(level) {
	case "debug":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
	case "info":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
	case "warn":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelWarn,
		}
	case "error":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelError,
		}
	default:
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
		err = ErrInvalidLogLevel
	}

	switch strings.ToLower(format) {
	case "text":
		slogHandler = slog.NewTextHandler(os.Stdout, &slogHandlerOptions)
	case "json":
		slogHandler = slog.NewJSONHandler(os.Stdout, &slogHandlerOptions)
	default:
		slogHandler = slog.NewTextHandler(os.Stdout, &slogHandlerOptions)
		err = ErrInvalidLogFormat
	}

	return slog.New(slogHandler), err
}
