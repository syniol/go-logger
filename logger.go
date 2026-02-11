// Copyright 2023-2026 Syniol Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Package sylog provides a simple way to create a searchable logs for your logging needs
//
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

type logLevel string

const (
	levelInfo      logLevel = "info"
	levelDebug     logLevel = "debug"
	levelNotice    logLevel = "notice"
	levelWarning   logLevel = "warn"
	levelAlert     logLevel = "alert"
	levelError     logLevel = "error"
	levelEmergency logLevel = "emergency"
	levelCritical  logLevel = "crit"
	levelFatal     logLevel = "fatal"
)

type logger struct {
	Level     logLevel      `json:"level"`
	Message   string        `json:"message"`
	Trace     []interface{} `json:"trace"`
	Timestamp string        `json:"timestamp"`
}

func logLocation() string {
	_, fileName, fileLine, ok := runtime.Caller(2)
	if !ok {
		return "file: path could not be found"
	}

	return fmt.Sprintf(
		"file: '%s' on line: %d",
		fileName,
		fileLine,
	)
}

func log(level logLevel, args []interface{}) *logger {
	logLocationDetail := logLocation()

	if len(args) == 0 {
		return &logger{
			Level:     level,
			Message:   "",
			Trace:     []interface{}{logLocationDetail},
			Timestamp: time.Now().Format(time.RFC3339),
		}
	}

	var allArgs []interface{}
	allArgs = append(allArgs, args...)

	allArgs = append(allArgs, logLocationDetail)

	return &logger{
		Level:   level,
		Message: allArgs[0].(string),
		Trace: func() []interface{} {
			return allArgs[1:]
		}(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func write(level logLevel, args []interface{}) {
	logContents, _ := json.Marshal(log(level, args))

	println(string(logContents))
}

// LogInfo logs messages where in output JSON key "level" is "info"
func LogInfo(args ...interface{}) {
	write(levelInfo, args)
}

// LogDebug logs messages where in output JSON key "level" is "debug"
func LogDebug(args ...interface{}) {
	write(levelDebug, args)
}

// LogNotice logs messages where in output JSON key "level" is "notice"
func LogNotice(args ...interface{}) {
	write(levelNotice, args)
}

// LogWarning logs messages where in output JSON key "level" is "warn"
func LogWarning(args ...interface{}) {
	write(levelWarning, args)
}

// LogAlert logs messages where in output JSON key "level" is "alert"
func LogAlert(args ...interface{}) {
	write(levelAlert, args)
}

// LogEmergency logs messages where in output JSON key "level" is "emergency"
func LogEmergency(args ...interface{}) {
	write(levelEmergency, args)
}

// LogError logs messages where in output JSON key "level" is "error"
func LogError(args ...interface{}) {
	write(levelError, args)
}

// LogCritical logs messages where in output JSON key "level" is "crit"
func LogCritical(args ...interface{}) {
	write(levelCritical, args)
}

// LogFatal logs messages where in output JSON key "level" is "fatal"
func LogFatal(args ...interface{}) {
	write(levelFatal, args)
}
