package autoheal

import (
	"testing"
)

func TestAutoHeal(t *testing.T) {
	containerName := []string{"ng"}
	result := AutoHeal(containerName)
	if (result != nil) {
		t.Errorf("Unsuccesful auto reload")
	}else {
		t.Logf("Succesfully autohealed")
	}
}