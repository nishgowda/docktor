package healthcheck 

import (
	"testing"
)

func TestPerform(t *testing.T) {
	var dummy = []string{"bdd7731b8121"}
	err := PerformHealthCheck(dummy)
	if (err != nil) {
		t.Errorf("Error in creating health check %s\n", err)
	}else {
		t.Logf("Succesfully created new container\n")
	}
}

