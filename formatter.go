package simplelogger

import (
	"fmt"
	"time"
)

// Formatter is responsible to handle the log format.
type Formatter interface {
	Format(level LogLevel, message string, data map[string]interface{}) (string, error)
}

// DefaultFormatter is the formatter for the DefaultLogger.
type DefaultFormatter struct{}

// Format returns the log with the date in RFC3339 format, loglevel, message and extra data as a string.
func (formatter DefaultFormatter) Format(level LogLevel, message string, data map[string]interface{}) (string, error) {
	formattedTime := time.Now().UTC().Format(time.RFC3339)
	if len(data) == 0 {
		return fmt.Sprintf("%s | [%s] | %s | No extra data\n", formattedTime, level, message), nil
	}
	return fmt.Sprintf("%s | [%s] | %s | %+v\n", formattedTime, level, message, data), nil
}
