package sender

import (
	"fmt"
	"testing"
)

func Test_DispatchTrainsIntegration(t *testing.T) {
	//figure out short
	if testing.Short() {
		t.Skip("integration test")
	}
	fmt.Println("checking if skipped")
	//call DispatchTrains
	//verify
}
