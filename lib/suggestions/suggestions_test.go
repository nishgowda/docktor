package suggestions

import (
	"testing"
)

func TestSuggestions(t *testing.T) {
	 filePath := "../../testdata/Dockerfile"
	 result := ReadImage(filePath)
	 if (result != nil) {
		 t.Errorf("Error in reading Dockerfile : %s\n", result)
	 } else {
		 t.Logf("Succesfully ran suggestions in Dockerfile\n")
	 }
}