package simplelogger

import "fmt"

// Writer is responsible to do any operation to display and/or store the formatted Log.
type Writer interface {
	Write(message string) error
}

// DefaultWriter is the Writer for the DefaultLogger.
type DefaultWriter struct{}

// Write displays the formatted Log message to the standard output.
func (writer DefaultWriter) Write(message string) error {
	fmt.Print(message)

	return nil
}
