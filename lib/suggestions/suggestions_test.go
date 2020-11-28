package suggestions

import (
	"testing"
)

func TestSuggestions(t *testing.T) {
	filePath := "../../testdata/Dockerfile"
	err := ReadImage(filePath)
	if err != nil {
		t.Errorf("Error in reading Dockerfile : %s\n", err)
	} else {
		t.Logf("Successfully ran suggestions in Dockerfile\n")
	}
}
