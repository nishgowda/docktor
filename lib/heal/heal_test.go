package heal

import (
	"testing"
)
func TestHeal(t *testing.T) {
	var containerIds = []string{"bdeb2a9b9bde"}
	result := ContainerHeal(containerIds)
	if (result != nil ){
		t.Errorf("Error in healing contiainer %s", result)
	}else{
		t.Logf("Succesfully healed contianer")
	}
}