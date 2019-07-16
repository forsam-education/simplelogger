package simplelogger

type LogLevel string

const (
	DEBUG    LogLevel = "DEBUG"
	INFO     LogLevel = "INFO"
	WARN     LogLevel = "WARN"
	ERROR    LogLevel = "ERROR"
	CRITICAL LogLevel = "CRITICAL"
)
