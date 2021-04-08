package heal

import (
	"testing"
)

func TestHeal(t *testing.T) {
	containers := GetUnheatlhyContainers()
	_, err := ContainerHeal(containers)
	if err != nil {
		t.Errorf("Error in healing container %s", err)
	} else {
		t.Logf("Successfully healed container")
	}
}
