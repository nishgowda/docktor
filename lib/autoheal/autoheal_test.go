package autoheal

import (
	"testing"
)

func TestAutoHeal(t *testing.T) {
	containerName := []string{"ng"}
	err := AutoHeal(containerName)
	if (err != nil) {
		t.Errorf("Unsuccesful auto reload")
	}else {
		t.Logf("Succesfully autohealed")
	}
}