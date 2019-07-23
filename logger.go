package simplelogger

import (
	"fmt"
)

// Logger is the basic struct of a logger.
type Logger struct {
	formatter Formatter
	writer    Writer
	MinLevel  LogLevel
}

// NewDefaultLogger returns a standard logger to StdOut using the default formatter.
func NewDefaultLogger() Logger {
	return Logger{DefaultFormatter{}, DefaultWriter{}, DEBUG}
}

func (logger Logger) Log(level LogLevel, message string, data LogExtraData) {
	if level < logger.MinLevel {
		return
	}

	formattedMessage, err := logger.formatter.Format(level, message, data)

	if err != nil {
		fmt.Printf("error while calling Log formatter: %s", err.Error())
		return
	}

	err = logger.writer.Write(formattedMessage)
	if err != nil {
		fmt.Printf("error while calling Log writer: %s", err.Error())
	}
}

// Debug handles logging of messages with the level Debug.
func (logger Logger) Debug(message string, data LogExtraData) {
	logger.Log(DEBUG, message, data)
}

// Info handles logging of messages with the level Info.
func (logger Logger) Info(message string, data LogExtraData) {
	logger.Log(INFO, message, data)
}

// Warn handles logging of messages with the level Warn.
func (logger Logger) Warn(message string, data LogExtraData) {
	logger.Log(WARN, message, data)
}

// Error handles logging of messages with the level Error.
func (logger Logger) Error(message string, data LogExtraData) {
	logger.Log(ERROR, message, data)
}

// Critical handles logging of messages with the level Critical.
func (logger Logger) Critical(message string, data LogExtraData) {
	logger.Log(CRITICAL, message, data)
}

// StdError handles logging objects of type 'error'.
func (logger Logger) StdError(err error, data LogExtraData) {
	logger.Log(ERROR, err.Error(), data)
}
