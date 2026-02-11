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
)

type logger struct {
	Level     logLevel      `json:"level"`
	Facility  string        `json:"facility"`
	Message   string        `json:"message"`
	Trace     []interface{} `json:"trace"`
	Timestamp string        `json:"timestamp"`
}

func logLocation() string {
	_, fileName, fileLine, ok := runtime.Caller(2)
	if !ok {
		return "location: path could not be found"
	}

	return fmt.Sprintf(
		"location: '%s' on line: %d",
		fileName,
		fileLine,
	)
}

func log(level logLevel, facility string, args []interface{}) *logger {
	logLocationDetail := logLocation()

	if len(args) == 0 {
		return &logger{
			Level:     level,
			Facility:  facility,
			Message:   "",
			Trace:     []interface{}{logLocationDetail},
			Timestamp: time.Now().Format(time.RFC3339),
		}
	}

	var allArgs []interface{}
	allArgs = append(allArgs, args...)

	allArgs = append(allArgs, logLocationDetail)

	return &logger{
		Level:    level,
		Facility: facility,
		Message:  allArgs[0].(string),
		Trace: func() []interface{} {
			return allArgs[1:]
		}(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func write(level logLevel, facility string, args []interface{}) {
	logContents, _ := json.Marshal(log(level, facility, args))

	println(string(logContents))
}

// LogInfo logs messages where in output JSON key "level" is "info"
func LogInfo(facility string, args ...interface{}) {
	write(levelInfo, facility, args)
}

// LogDebug logs messages where in output JSON key "level" is "debug"
func LogDebug(facility string, args ...interface{}) {
	write(levelDebug, facility, args)
}

// LogNotice logs messages where in output JSON key "level" is "notice"
func LogNotice(facility string, args ...interface{}) {
	write(levelNotice, facility, args)
}

// LogWarning logs messages where in output JSON key "level" is "warn"
func LogWarning(facility string, args ...interface{}) {
	write(levelWarning, facility, args)
}

// LogAlert logs messages where in output JSON key "level" is "alert"
func LogAlert(facility string, args ...interface{}) {
	write(levelAlert, facility, args)
}

// LogEmergency logs messages where in output JSON key "level" is "emergency"
func LogEmergency(facility string, args ...interface{}) {
	write(levelEmergency, facility, args)
}

// LogError logs messages where in output JSON key "level" is "error"
func LogError(facility string, args ...interface{}) {
	write(levelError, facility, args)
}

// LogCritical logs messages where in output JSON key "level" is "crit"
func LogCritical(facility string, args ...interface{}) {
	write(levelCritical, facility, args)
}
