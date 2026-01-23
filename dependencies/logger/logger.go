package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type AppLogger struct {
	logger *slog.Logger
}

// Debug logs at debug level
func (l AppLogger) Debug(args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelDebug, fmt.Sprint(args...))
}

func (l AppLogger) Debugf(format string, args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelDebug, fmt.Sprintf(format, args...))
}

func (l AppLogger) Info(args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelInfo, fmt.Sprint(args...))
}

func (l AppLogger) Infof(format string, args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelInfo, fmt.Sprintf(format, args...))
}

func (l AppLogger) Error(args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelError, fmt.Sprint(args...))
}

func (l AppLogger) Errorf(format string, args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelError, fmt.Sprintf(format, args...))
}

func (l AppLogger) Warning(args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelWarn, fmt.Sprint(args...))
}

func (l AppLogger) Warningf(format string, args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelWarn, fmt.Sprintf(format, args...))
}

func (l AppLogger) WithError(err error, message string) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelError, message, slog.Any("error", err))
}

func (l AppLogger) WithErrorf(err error, format string, args ...interface{}) {
	l.logger.LogAttrs(contextOrBackground(), slog.LevelError, fmt.Sprintf(format, args...), slog.Any("error", err))
}

// NewTextLogger returns a human-friendly console logger
func NewTextLogger() AppLogger {
	level := parseLevel(os.Getenv("LOG_LEVEL"))
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	l := slog.New(h).With(defaultAttrs()...)
	return AppLogger{logger: l}
}

// NewJsonLogger returns a JSON logger suitable for structured logs
func NewJsonLogger() AppLogger {
	level := parseLevel(os.Getenv("LOG_LEVEL"))
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	l := slog.New(h).With(defaultAttrs()...)
	return AppLogger{logger: l}
}

func parseLevel(val string) slog.Leveler {
	switch val {
	case "debug", "DEBUG":
		return slog.LevelDebug
	case "warn", "warning", "WARN":
		return slog.LevelWarn
	case "error", "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func defaultAttrs() []any {
	env := os.Getenv("DD_ENV")
	if env == "" {
		env = "dev"
	}
	return []any{
		slog.String("env", env),
		slog.Int("data_privacy", 0),
	}
}

// contextOrBackground provides a context without leaking external deps
func contextOrBackground() context.Context {
	return context.Background()
}
