package simplelogger

// LogLevel represents a... log level
type LogLevel uint

const (
	// DEBUG log level
	DEBUG LogLevel = iota
	// INFO log level
	INFO LogLevel = iota
	// WARN log level
	WARN LogLevel = iota
	// ERROR log level
	ERROR LogLevel = iota
	// CRITICAL log level
	CRITICAL LogLevel = iota
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "CRITICAL"}[l]
}
