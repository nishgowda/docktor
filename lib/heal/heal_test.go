package heal

import (
	"testing"
)
func TestHeal(t *testing.T) {
	containers := GetUnheatlhyContainers()
	err := ContainerHeal(containers)
	if (err != nil ){
		t.Errorf("Error in healing contiainer %s", err)
	}else{
		t.Logf("Succesfully healed contianer")
	}
}