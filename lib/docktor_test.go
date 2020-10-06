package lib 

import (
	"testing"
)

func TestContainer(t *testing.T) {
	result := CreateContainer("registry.gitlab.com/tweetcode/back-main:1.0", "0.0.0.0", "4000")
	if (result != "Success") {
		t.Errorf("Error in creating health check")
	}else {
		t.Logf("Succesfully created new container")
	}
}