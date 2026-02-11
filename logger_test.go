package sylog

import (
	"testing"
)

func TestLogWithEmptyArgument(t *testing.T) {
	LogInfo("PayNow Microservice")
	LogEmergency("PayNow Microservice")
	LogCritical("PayNow Microservice")
	LogError("PayNow Microservice")
	LogDebug("PayNow Microservice")
	LogAlert("PayNow Microservice")
	LogNotice("PayNow Microservice")
}

func TestLog(t *testing.T) {
	LogInfo("PayNow Microservice", "SomeText", "lorem ipsum", struct {
		Category string
	}{
		Category: "MyCategory",
	}, 3232)

	LogWarning("PayNow Microservice", "SomeText", "lorem ipsum", struct {
		Category string
	}{
		Category: "MyCategory",
	}, 3232)

	LogEmergency("PayNow Microservice", "SomeText")
	LogCritical("PayNow Microservice", "SomeText")
	LogError("PayNow Microservice", "SomeText")
	LogDebug("PayNow Microservice", "SomeText")
	LogAlert("PayNow Microservice", "SomeText")
	LogNotice("PayNow Microservice", "SomeText")
}
