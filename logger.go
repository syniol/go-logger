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
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
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
	Level     logLevel `json:"level"`
	Facility  string   `json:"facility"`
	Message   string   `json:"message"`
	Trace     []string `json:"trace"`
	Timestamp string   `json:"timestamp"`
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

func log(level logLevel, facility string, args []string) *logger {
	logLocationDetail := logLocation()

	if len(args) == 0 {
		return &logger{
			Level:     level,
			Facility:  facility,
			Message:   "",
			Trace:     []string{logLocationDetail},
			Timestamp: time.Now().Format(time.RFC3339),
		}
	}

	var allArgs []string
	allArgs = append(allArgs, args...)

	allArgs = append(allArgs, logLocationDetail)

	return &logger{
		Level:    level,
		Facility: facility,
		Message:  allArgs[0],
		Trace: func() []string {
			return allArgs[1:]
		}(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// bufferPool recycles *bytes.Buffer instances to minimize heap allocations.
var bufferPool = sync.Pool{
	New: func() interface{} {
		// Pre-allocate 512 bytes to handle most log lines without re-allocating.
		return bytes.NewBuffer(make([]byte, 0, 512))
	},
}

var writer io.Writer = os.Stdout

// setOutput allows the user (or benchmark) to change where logs are sent.
func setOutput(w io.Writer) {
	writer = w
}

func write(level logLevel, facility string, args []string) {
	loggerData := log(level, facility, args)

	// Snag a buffer from the pool (thread-safe).
	buf := bufferPool.Get().(*bytes.Buffer)

	// Ensure it's returned to the pool after the function exits.
	defer func() {
		// PRO TIP: If a log was massive (e.g. 1MB), don't put it back in the pool.
		// This prevents "Memory Bloat" where the pool holds huge unused chunks.
		if buf.Cap() <= 64*1024 { // 64KB Limit
			buf.Reset()
			bufferPool.Put(buf)
		}
	}()

	// Build the JSON manually to avoid json.Marshal (Reflection = Slow).
	buf.WriteString(`{"level":"`)
	buf.WriteString(string(loggerData.Level))
	buf.WriteString(`","facility":"`)
	buf.WriteString(facility)
	buf.WriteString(`","message":"`)
	buf.WriteString(loggerData.Message)
	buf.WriteString(`","trace":[`)
	for i, trace := range loggerData.Trace {
		buf.WriteString(`"` + trace + `"`)
		if i < len(loggerData.Trace)-1 {
			buf.WriteString(`,`)
		}
	}
	buf.WriteString(`],"timestamp":"`)
	buf.WriteString(loggerData.Timestamp)

	buf.WriteString(`"}`)
	buf.WriteByte('\n')

	_, _ = io.Copy(writer, buf)

	setOutput(writer)
	_, _ = writer.Write(buf.Bytes())
}

// LogInfo logs messages where in output JSON key "level" is "info"
func LogInfo(facility string, args ...string) {
	write(levelInfo, facility, args)
}

// LogDebug logs messages where in output JSON key "level" is "debug"
func LogDebug(facility string, args ...string) {
	write(levelDebug, facility, args)
}

// LogNotice logs messages where in output JSON key "level" is "notice"
func LogNotice(facility string, args ...string) {
	write(levelNotice, facility, args)
}

// LogWarning logs messages where in output JSON key "level" is "warn"
func LogWarning(facility string, args ...string) {
	write(levelWarning, facility, args)
}

// LogAlert logs messages where in output JSON key "level" is "alert"
func LogAlert(facility string, args ...string) {
	write(levelAlert, facility, args)
}

// LogEmergency logs messages where in output JSON key "level" is "emergency"
func LogEmergency(facility string, args ...string) {
	write(levelEmergency, facility, args)
}

// LogError logs messages where in output JSON key "level" is "error"
func LogError(facility string, args ...string) {
	write(levelError, facility, args)
}

// LogCritical logs messages where in output JSON key "level" is "crit"
func LogCritical(facility string, args ...string) {
	write(levelCritical, facility, args)
}
