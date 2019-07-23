package simplelogger

// LogLevel represents a... Log level
type LogLevel uint

const (
	// DEBUG Log level
	DEBUG LogLevel = iota
	// INFO Log level
	INFO LogLevel = iota
	// WARN Log level
	WARN LogLevel = iota
	// ERROR Log level
	ERROR LogLevel = iota
	// CRITICAL Log level
	CRITICAL LogLevel = iota
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "CRITICAL"}[l]
}
