package simplelogger

import (
	"fmt"
	"time"
)

// Formatter is responsible to handle the Log format.
type Formatter interface {
	Format(level LogLevel, message string, data LogExtraData) (string, error)
}

// LogExtraData is the basic structure to pass extra data to a Log.
type LogExtraData map[string]interface{}

// DefaultFormatter is the Formatter for the DefaultLogger.
type DefaultFormatter struct{}

// Format returns the Log with the date in RFC3339 format, loglevel, message and extra data as a string.
func (formatter DefaultFormatter) Format(level LogLevel, message string, data LogExtraData) (string, error) {
	formattedTime := time.Now().UTC().Format(time.RFC3339)
	if len(data) == 0 {
		return fmt.Sprintf("%s | [%s] | %s\n", formattedTime, level.String(), message), nil
	}
	return fmt.Sprintf("%s | [%s] | %s | %+v\n", formattedTime, level, message, data), nil
}
