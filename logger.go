// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// Package sylog provides a simple way to create a searchable logs for your logging needs

//	package main
//
//	import (
//		"github.com/syniol/go-logger"
//	)
//
//	func main() {
//		sylog.LogAlert("mock message for alert log")
//	}

package sylog

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type LogLevel string

const (
	levelInfo      LogLevel = "info"
	levelDebug     LogLevel = "debug"
	levelNotice    LogLevel = "notice"
	levelWarning   LogLevel = "warn"
	levelAlert     LogLevel = "alert"
	levelError     LogLevel = "error"
	levelEmergency LogLevel = "emergency"
	levelCritical  LogLevel = "crit"
	levelFatal     LogLevel = "fatal"
)

type logger struct {
	Level     LogLevel      `json:"level"`
	Message   string        `json:"message"`
	Trace     []interface{} `json:"trace"`
	Timestamp string        `json:"timestamp"`
}

func log(level LogLevel, args []interface{}) *logger {
	if len(args) == 0 {
		return &logger{
			Level:     level,
			Message:   "empty message",
			Trace:     nil,
			Timestamp: time.Now().Format(time.RFC3339),
		}
	}

	var allArgs []interface{}
	allArgs = append(allArgs, args...)

	_, fileName, fileLine, _ := runtime.Caller(2)

	allArgs = append(allArgs, fmt.Sprintf(
		"location: file: '%s' on line: %d",
		fileName,
		fileLine,
	))

	return &logger{
		Level:   level,
		Message: allArgs[0].(string),
		Trace: func() []interface{} {
			return allArgs[1:]
		}(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func write(level LogLevel, args []interface{}) {
	logContents, _ := json.Marshal(log(level, args))

	println(string(logContents))
}

func LogInfo(args ...interface{}) {
	write(levelInfo, args)
}

func LogDebug(args ...interface{}) {
	write(levelDebug, args)
}

func LogNotice(args ...interface{}) {
	write(levelNotice, args)
}

func LogWarning(args ...interface{}) {
	write(levelWarning, args)
}

func LogAlert(args ...interface{}) {
	write(levelAlert, args)
}

func LogEmergency(args ...interface{}) {
	write(levelEmergency, args)
}

func LogError(args ...interface{}) {
	write(levelError, args)
}

func LogCritical(args ...interface{}) {
	write(levelCritical, args)
}

func LogFatal(args ...interface{}) {
	write(levelFatal, args)
}
