package healthcheck 

import (
	"testing"
)

func TestPerform(t *testing.T) {
	var dummy = []string{"bdd7731b8121"}
	result := PerformHealthCheck(dummy)
	if (result != nil) {
		t.Errorf("Error in creating health check: %s", result)
	}else {
		t.Logf("Succesfully created new container")
	}
}

