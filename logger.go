package simplelogger

import (
	"fmt"
)

type Logger struct {
	formatter Formatter
	writer    Writer
}

func NewDefaultLogger() Logger {
	return Logger{DefaultFormatter{}, DefaultWriter{}}
}

func (logger Logger) Log(level LogLevel, message string, data map[string]interface{}) {
	formattedMessage, err := logger.formatter.Format(level, message, data)

	if err != nil {
		fmt.Printf("error while calling log formatter: %s", err.Error())
		return
	}

	err = logger.writer.Write(formattedMessage)
	if err != nil {
		fmt.Printf("error while calling log writer: %s", err.Error())
	}
}

func (logger Logger) Debug(message string, data map[string]interface{}) {
	logger.Log(DEBUG, message, data)
}

func (logger Logger) Info(message string, data map[string]interface{}) {
	logger.Log(INFO, message, data)
}

func (logger Logger) Warn(message string, data map[string]interface{}) {
	logger.Log(WARN, message, data)
}

func (logger Logger) Error(message string, data map[string]interface{}) {
	logger.Log(ERROR, message, data)
}

func (logger Logger) Critical(message string, data map[string]interface{}) {
	logger.Log(CRITICAL, message, data)
}

func (logger Logger) StdError(err error, data map[string]interface{}) {
	logger.Log(ERROR, err.Error(), data)
}
