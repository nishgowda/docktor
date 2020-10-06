package lib 

import (
	"testing"
)

func TestContainer(t *testing.T) {
	result := Perform()
	if (result != nil) {
		t.Errorf("Error in creating health check: %s", result)
	}else {
		t.Logf("Succesfully created new container")
	}
}