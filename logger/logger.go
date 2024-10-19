package logger

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

// Log level constants
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var logLevel = INFO // Default log level

// SetLogLevel sets the global log level
func SetLogLevel(level int) {
	logLevel = level
}
func ConvertStringToLogLevel(level string) int {
	// make the level uppercase
	level = strings.ToUpper(level)
	// Map the string log level to an integer
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	default:
		Error("Invalid log level. Use: DEBUG, INFO, WARN, ERROR")
		os.Exit(-1)
	}
	return -1
}

// logMessage is the internal logging function that checks log levels
func logMessage(level int, prefix string, v ...interface{}) {
	if level >= logLevel {
		// Use log.Println to handle multiple arguments like fmt.Println
		log.Println(append([]interface{}{prefix}, v...)...)
	}
}

// Debug logs a message at the DEBUG level
func Debug(v ...interface{}) {
	logMessage(DEBUG, "[DEBUG]", v...)
}

// Info logs a message at the INFO level
func Info(v ...interface{}) {
	logMessage(INFO, "[INFO]", v...)
}

// Warn logs a message at the WARN level
func Warn(v ...interface{}) {
	logMessage(WARN, "[WARN]", v...)
}

// Error logs a message at the ERROR level
func Error(v ...interface{}) {
	logMessage(ERROR, "[ERROR]", v...)
}

func LogStackTraceAndExit(err interface{}) {
	// log erro
	if err != nil {
		Error(err)
	}
	// log stack trace
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	Error("Stack trace:\n", string(buf))
	os.Exit(-1)
}

// Route logs to whichever file
func SetOutput(w io.Writer) {
	log.SetOutput(w)
}
