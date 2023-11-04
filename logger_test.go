package sylog

import (
	"testing"
)

func TestLog(t *testing.T) {
	LogInfo("SomeText", "lorem ipsum", struct {
		Category string
	}{
		Category: "Categosdsds",
	}, 3232)

	LogWarning("SomeText", "lorem ipsum", struct {
		Category string
	}{
		Category: "Categosdsds",
	}, 3232)

	LogEmergency("SomeText")
	LogCritical("SomeText")
	LogFatal("SomeText")
	LogError("SomeText")
	LogDebug("SomeText")
	LogAlert("SomeText")
	LogNotice("SomeText")
}
