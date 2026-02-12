package sylog

import "testing"

func TestNew(t *testing.T) {
	newLogger := New("pay-service")

	newLogger.LogInfo("payment has been completed successfully with ID: 2322-1321-3211")
}
