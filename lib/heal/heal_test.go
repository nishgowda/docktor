package heal

import (
	"testing"
)
func TestHeal(t *testing.T) {
	containers := GetUnheatlhyContainers()
	result := ContainerHeal(containers)
	if (result != nil ){
		t.Errorf("Error in healing contiainer %s", result)
	}else{
		t.Logf("Succesfully healed contianer")
	}
}