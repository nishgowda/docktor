package healthcheck

import (
	"testing"
)

func TestPerform(t *testing.T) {
	var dummy = []string{"bdd7731b8121"}
	_, err := PerformHealthCheck(dummy)
	if err != nil {
		t.Errorf("Error in creating health check %s\n", err)
	} else {
		t.Logf("Successfully created new container\n")
	}
}
