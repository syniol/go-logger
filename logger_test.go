package sylog

import (
	"encoding/json"
	"errors"
	"testing"
	"time"
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

func TestErrorMarshall(t *testing.T) {
	level := levelCritical
	facility := "unknown"
	err := errors.New("JSON Marshall error")

	errMessage := `{"level":"` +
		string(level) +
		`", "facility":"` +
		facility +
		`", "message":"error creating a log - ` +
		err.Error() +
		`", "trace":[]` +
		`, "timestamp":"` +
		time.Now().Format(time.RFC3339) +
		`" }`

	var handMadeLog logger
	marshallErr := json.Unmarshal([]byte(errMessage), &handMadeLog)
	if marshallErr != nil {
		t.Error(marshallErr)
	}

	if handMadeLog.Level != level {
		t.Errorf("was expecting %s", handMadeLog.Level)
	}
}
