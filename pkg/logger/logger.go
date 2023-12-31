// Package logger provides a logging functionality.
package logger

import (
	"context"

	// third party
	"go.uber.org/zap"
)

// Logger Interface - represents logger interface.
type Logger interface {
	Named(name string) Logger
	With(args ...interface{}) Logger
	WithContext(ctx context.Context) Logger
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
	Unwrap() *zap.Logger
}
