# Go Simple Logger

[![CircleCI](https://circleci.com/gh/forsam-education/simplelogger.svg?style=svg)](https://circleci.com/gh/forsam-education/simplelogger)
[![GoDoc](https://godoc.org/github.com/forsam-education/simplelogger?status.svg)](https://godoc.org/github.com/forsam-education/simplelogger)
[![Go Report Card](https://goreportcard.com/badge/github.com/forsam-education/simplelogger)](https://goreportcard.com/report/github.com/forsam-education/simplelogger)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fforsam-education%2Fsimplelogger.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fforsam-education%2Fsimplelogger?ref=badge_shield)

## Background

While writing some lambdas in Golang, I found myself searching for a very light, adaptable logger that would fit for basic printing, elasticsearch indexing or any other logging system, with the most simple interface possible.

The result is this logger.

## How to implement

This package exports a Logger struct defined like this (you'll find an interface representation just bellow):
```go
package simplelogger

type LogExtraData map[string]interface{}

type Logger struct {
	formatter Formatter
	writer    Writer
	MinLevel  LogLevel
}

type _ interface {
	Log(level LogLevel, message string, data LogExtraData)
	Debug(message string, data LogExtraData)
	Info(message string, data LogExtraData)
	Warn(message string, data LogExtraData)
	Error(message string, data LogExtraData)
	Critical(message string, data LogExtraData)
	StdErrorCritical(err error, data LogExtraData)
	StdError(err error, data LogExtraData)
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

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fforsam-education%2Fsimplelogger.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fforsam-education%2Fsimplelogger?ref=badge_large)