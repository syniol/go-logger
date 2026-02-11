package sylog

import (
	"testing"
)

func TestLogWithEmptyArgument(t *testing.T) {
	LogInfo()
	LogEmergency()
	LogCritical()
	LogFatal()
	LogError()
	LogDebug()
	LogAlert()
	LogNotice()
}

func TestLog(t *testing.T) {
	LogInfo("SomeText", "lorem ipsum", struct {
		Category string
	}{
		Category: "MyCategory",
	}, 3232)

	LogWarning("SomeText", "lorem ipsum", struct {
		Category string
	}{
		Category: "MyCategory",
	}, 3232)

	LogEmergency("SomeText")
	LogCritical("SomeText")
	LogFatal("SomeText")
	LogError("SomeText")
	LogDebug("SomeText")
	LogAlert("SomeText")
	LogNotice("SomeText")
}
