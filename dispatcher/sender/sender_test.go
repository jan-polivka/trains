package sender

import (
	"testing"
)

func Test_DispatchTrainsIntegration(t *testing.T) {
	//figure out short
	if testing.Short() {
		t.Skip("integration test")
	}
	//call DispatchTrains
	//verify
}
