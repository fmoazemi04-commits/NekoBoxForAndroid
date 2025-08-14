package neko_log

import (
	"io"
	"os"
	"sync"
)

// LogWriter interface for log writing
type LogWriterInterface interface {
	Truncate() error
	Write([]byte) (int, error)
}

// SimpleLogWriter implements basic file logging
type SimpleLogWriter struct {
	file   *os.File
	mutex  sync.Mutex
	maxSize int
}

// Global log writer instance
var LogWriter LogWriterInterface

// Global settings
var (
	LogWriterDisable bool
	TruncateOnStart  bool
)

// SetupLog initializes the logging system
func SetupLog(maxSizeBytes int, logPath string) error {
	writer := &SimpleLogWriter{
		maxSize: maxSizeBytes,
	}
	
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	
	writer.file = file
	LogWriter = writer
	
	if TruncateOnStart {
		LogWriter.Truncate()
	}
	
	return nil
}

// Truncate clears the log file
func (w *SimpleLogWriter) Truncate() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	
	if w.file != nil {
		return w.file.Truncate(0)
	}
	return nil
}

// Write writes data to the log file
func (w *SimpleLogWriter) Write(data []byte) (int, error) {
	if LogWriterDisable {
		return len(data), nil
	}
	
	w.mutex.Lock()
	defer w.mutex.Unlock()
	
	if w.file == nil {
		return 0, io.ErrClosedPipe
	}
	
	// Check file size and truncate if necessary
	if stat, err := w.file.Stat(); err == nil {
		if stat.Size() > int64(w.maxSize) {
			w.file.Truncate(0)
			w.file.Seek(0, 0)
		}
	}
	
	return w.file.Write(data)
}