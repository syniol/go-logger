package sylog

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestLogWithEmptyArgument(t *testing.T) {
	LogInfo("PayNow Microservice", "Mocked Message")
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
	timestamp := time.Now().Format(time.RFC3339)

	errMessage := `{"level":"` +
		string(level) +
		`", "facility":"` +
		facility +
		`", "message":"error creating a log - ` +
		err.Error() +
		`", "trace":["` + logLocation() + `"]` +
		`, "timestamp":"` +
		timestamp +
		`" }`

	var handMadeLog logger
	marshallErr := json.Unmarshal([]byte(errMessage), &handMadeLog)
	if marshallErr != nil {
		t.Error(marshallErr)
	}

	if handMadeLog.Level != level {
		t.Errorf("was expecting %s", level)
	}

	if handMadeLog.Facility != facility {
		t.Errorf("was expecting %s", facility)
	}

	if !strings.Contains(handMadeLog.Message, err.Error()) {
		t.Errorf("was expecting error to contain %s", err.Error())
	}

	if handMadeLog.Timestamp != timestamp {
		t.Errorf("was expecting %s", timestamp)
	}

	if len(handMadeLog.Trace) != 1 && !strings.Contains(handMadeLog.Trace[0].(string), "location: ") {
		t.Errorf("was expecting 1 trace for file location")
	}
}
