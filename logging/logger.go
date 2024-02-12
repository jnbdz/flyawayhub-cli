package logging

import (
	"go.uber.org/zap"
)

// Logger defines a generic logging interface with methods for various logging levels.
type Logger interface {
	// Debug logs are typically voluminous, and are usually disabled in production.
	Debug(msg string, keysAndValues ...interface{})

	// Info is for general operational entries about what's happening inside the application.
	Info(msg string, keysAndValues ...interface{})

	// Warn is for logging messages about potentially harmful situations.
	Warn(msg string, keysAndValues ...interface{})

	// Error is for logging error messages that might still allow the application to continue running.
	Error(msg string, keysAndValues ...interface{})

	// DPanic logs messages at the Panic level but doesn’t panic in the production environment.
	DPanic(msg string, keysAndValues ...interface{})

	// Panic is for logging messages that indicate situations that should never occur.
	Panic(msg string, keysAndValues ...interface{})

	// Fatal is for logging fatal messages. The system shuts down after logging the message.
	Fatal(msg string, keysAndValues ...interface{})
}

var logger Logger

// zapLogger implements Logger with a Zap logger
type zapLogger struct {
	logger *zap.SugaredLogger
}

// NewZapLogger creates and returns a new Logger implemented using Zap
func NewZapLogger() Logger {
	logger, _ := zap.NewProduction()
	sugaredLogger := logger.Sugar()
	return &zapLogger{logger: sugaredLogger}
}

func SetLogger(newLogger Logger) {
	logger = newLogger
}

func ResetLogger() {
	logger = NewZapLogger()
}

func (z *zapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logger.Infow(msg, keysAndValues...)
}

func (z *zapLogger) Error(msg string, keysAndValues ...interface{}) {
	z.logger.Errorw(msg, keysAndValues...)
}

func (z *zapLogger) Debug(msg string, keysAndValues ...interface{}) {
	z.logger.Debugw(msg, keysAndValues...)
}

func (z *zapLogger) Warn(msg string, keysAndValues ...interface{}) {
	z.logger.Warnw(msg, keysAndValues...)
}

func (z *zapLogger) DPanic(msg string, keysAndValues ...interface{}) {
	z.logger.DPanicw(msg, keysAndValues...)
}

func (z *zapLogger) Panic(msg string, keysAndValues ...interface{}) {
	z.logger.Panicw(msg, keysAndValues...)
}

func (z *zapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	z.logger.Fatalw(msg, keysAndValues...)
}
