# Go Simple Logger

[![CircleCI](https://circleci.com/gh/forsam-education/simplelogger.svg?style=svg)](https://circleci.com/gh/forsam-education/simplelogger)
[![GoDoc](https://godoc.org/github.com/forsam-education/simplelogger?status.svg)](https://godoc.org/github.com/forsam-education/simplelogger)
[![Go Report Card](https://goreportcard.com/badge/github.com/forsam-education/simplelogger)](https://goreportcard.com/report/github.com/forsam-education/simplelogger)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Background

While writing some lambdas in Golang, I found myself searching for a very light, adaptable logger that would fit for basic printing, elasticsearch indexing or any other logging system, with the most simple interface possible.

The result is this logger.

## How to implement

This package exports a Logger struct defined like this (you'll find an interface representation just bellow):
```go
package simplelogger

type Logger struct {
	formatter Formatter
	writer    Writer
}

type _ interface {
	log(level LogLevel, message string, data map[string]interface{})
	Debug(message string, data map[string]interface{})
	Info(message string, data map[string]interface{})
	Warn(message string, data map[string]interface{})
	Error(message string, data map[string]interface{})
	Critical(message string, data map[string]interface{})
	StdError(err error, data map[string]interface{})
}
```

You can either use the DefaultLogger that logs to standard output with this format:

`2019-07-16T17:07:46Z | [ERROR] | issou l'erreur | map[extra_data:data]`

this way: 
```go
package yourpackage

logger := simplelogger.NewDefaultLogger()
```

Or you can implement your own writer and formatter according to these interfaces:

```go
package simplelogger

type Formatter interface {
	Format(level LogLevel, message string, data map[string]interface{}) (string, error)
}

type Writer interface {
	Write(message string) error
}
```

## Other Go Loggers

This is of course not an exhaustive list but here are some logger you may want to check out!

- [apsdehal/go-logger](https://github.com/apsdehal/go-logger)
- [azer/logger](https://github.com/azer/logger)