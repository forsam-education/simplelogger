package simplelogger

import (
	"fmt"
	"time"
)

type Formatter interface {
	Format(level LogLevel, message string, data map[string]interface{}) (string, error)
}

type DefaultFormatter struct{}

func (formatter DefaultFormatter) Format(level LogLevel, message string, data map[string]interface{}) (string, error) {
	formattedTime := time.Now().UTC().Format(time.RFC3339)
	if len(data) == 0 {
		return fmt.Sprintf("%s | [%s] | %s | No extra data\n", formattedTime, level, message), nil
	}
	return fmt.Sprintf("%s | [%s] | %s | %+v\n", formattedTime, level, message, data), nil
}
