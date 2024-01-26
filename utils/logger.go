package utils

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	InfoLevel LogLevel = iota
	DebugLevel
	ErrorLevel
)

var logLevelMapping = map[LogLevel]string{
	InfoLevel:  "INFO",
	DebugLevel: "DEBUG",
	ErrorLevel: "ERROR",
}

var logger = log.New(os.Stdout, "", log.LstdFlags)

// SetLogLevel sets the log level for the logger
func SetLogLevel(level LogLevel) {
	logger.SetPrefix(logLevelMapping[level] + ": ")
}

// Info logs an info message
func Info(message string) {
	logMessage(InfoLevel, message)
}

// Debug logs a debug message
func Debug(message string) {
	logMessage(DebugLevel, message)
}

// Error logs an error message
func Error(message string) {
	logMessage(ErrorLevel, message)
}

func logMessage(level LogLevel, message string) {
	logger.Println(fmt.Sprintf("[%s] %s", logLevelMapping[level], message))
}
