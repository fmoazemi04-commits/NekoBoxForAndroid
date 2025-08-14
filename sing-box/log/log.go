package log

import (
	"log"
)

// Logger represents a logger interface
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

// DefaultLogger implements the Logger interface
type DefaultLogger struct{}

// Debug logs debug messages
func (l *DefaultLogger) Debug(args ...interface{}) {
	log.Print("[DEBUG] ", args)
}

// Info logs info messages
func (l *DefaultLogger) Info(args ...interface{}) {
	log.Print("[INFO] ", args)
}

// Warn logs warning messages
func (l *DefaultLogger) Warn(args ...interface{}) {
	log.Print("[WARN] ", args)
}

// Error logs error messages
func (l *DefaultLogger) Error(args ...interface{}) {
	log.Print("[ERROR] ", args)
}

// NewDefaultLogger creates a new default logger
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// PlatformWriter represents a platform-specific log writer
type PlatformWriter interface {
	DisableColors() bool
	WriteMessage(level uint8, message string)
}

// DefaultPlatformWriter implements PlatformWriter
type DefaultPlatformWriter struct {
	// Implementation details
}

// DisableColors returns whether colors should be disabled
func (w *DefaultPlatformWriter) DisableColors() bool {
	return true
}

// WriteMessage writes a log message
func (w *DefaultPlatformWriter) WriteMessage(level uint8, message string) {
	log.Printf("[%d] %s", level, message)
}

// NewPlatformWriter creates a new platform writer
func NewPlatformWriter() PlatformWriter {
	return &DefaultPlatformWriter{}
}