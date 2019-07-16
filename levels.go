package simplelogger

// LogLevel represents a... log level
type LogLevel string

const (
	// DEBUG log level
	DEBUG LogLevel = "DEBUG"
	// INFO log level
	INFO LogLevel = "INFO"
	// WARN log level
	WARN LogLevel = "WARN"
	// ERROR log level
	ERROR LogLevel = "ERROR"
	// CRITICAL log level
	CRITICAL LogLevel = "CRITICAL"
)
