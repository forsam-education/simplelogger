package gologger

import "fmt"

type Writer interface {
	Write(message string) error
}

type DefaultWriter struct{}

func (writer DefaultWriter) Write(message string) error {
	fmt.Print(message)

	return nil
}
