package sylog

import "testing"

func TestNew(t *testing.T) {
	newLogger := NewLogger("pay-service", nil)

	newLogger.LogInfo("payment has been completed successfully with ID: 2322-1321-3211")
}
