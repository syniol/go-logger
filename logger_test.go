package sylog

import (
	"bytes"
	"strings"
	"testing"
)

func TestLoggedMessage(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	LogInfo("test message")

	output := buf.String()
	if !strings.Contains(output, "test message") {
		t.Errorf("Expected log to contain 'test message', got: %s", output)
	}
}

func TestLogWithEmptyArgumentExecution(t *testing.T) {
	LogInfo("PayNow Microservice", "Mocked Message")
	LogEmergency("PayNow Microservice")
	LogCritical("PayNow Microservice")
	LogError("PayNow Microservice")
	LogDebug("PayNow Microservice")
	LogAlert("PayNow Microservice")
	LogNotice("PayNow Microservice")
}

func TestLogExecution(t *testing.T) {
	LogInfo("PayNow Microservice", "SomeText", "lorem ipsum", " 3232")

	LogWarning("PayNow Microservice", "SomeText", "lorem ipsum", "3232")

	LogEmergency("PayNow Microservice", "SomeText")
	LogCritical("PayNow Microservice", "SomeText")
	LogError("PayNow Microservice", "SomeText")
	LogDebug("PayNow Microservice", "SomeText")
	LogAlert("PayNow Microservice", "SomeText")
	LogNotice("PayNow Microservice", "SomeText")
}
