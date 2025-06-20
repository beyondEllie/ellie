package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/tacheraSasi/ellie/configs"
)

const (
	LOG_FILENAME = "ellie.combined.log"
	CONFIG_DIR   = "/"+configs.ConfigDirName
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
	once          sync.Once
)

// initializes loggers with both file and console output
func initialize() {
	once.Do(func() {
		// Creates logs directory if it doesn't exist
		logDir := filepath.Join(CONFIG_DIR, "logs")
		if err := os.MkdirAll(logDir, 0755); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		// Opens log file
		logFile, err := os.OpenFile(
			filepath.Join(logDir, LOG_FILENAME),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0644,
		)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		// Creates multiwriter for both file and stdout
		multiWriter := io.MultiWriter(os.Stdout, logFile)

		// Initializes loggers with different prefixes and flags
		infoLogger = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		warningLogger = log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		errorLogger = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		debugLogger = log.New(multiWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

// Info logs an informational message
func Info(format string, v ...any) {
	initialize()
	infoLogger.Printf(format, v...)
}

// Warning logs a warning message
func Warning(format string, v ...any) {
	initialize()
	warningLogger.Printf(format, v...)
}

// Error logs an error message
func Error(err error) {
	initialize()
	if err != nil {
		errorLogger.Printf("%v", err)
	}
}

// Errorf logs a formatted error message
func Errorf(format string, v ...any) {
	initialize()
	errorLogger.Printf(format, v...)
}

// Debug logs a debug message
func Debug(format string, v ...any) {
	initialize()
	debugLogger.Printf(format, v...)
}

// Fatal logs a fatal error and exits the program
func Fatal(err error) {
	initialize()
	if err != nil {
		errorLogger.Fatalf("%v", err)
	}
}

// Fatalf logs a formatted fatal error and exits the program
func Fatalf(format string, v ...any) {
	initialize()
	errorLogger.Fatalf(format, v...)
}
