package simplelogger

import "fmt"

// Writer is responsible to do any operation to display and/or store the formatted log.
type Writer interface {
	Write(message string) error
}

// DefaultWriter is the writer for the DefaultLogger.
type DefaultWriter struct{}

// Write displays the formatted log message to the standard output.
func (writer DefaultWriter) Write(message string) error {
	fmt.Print(message)

	return nil
}
