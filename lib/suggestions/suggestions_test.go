package suggestions

import (
	"testing"
)

func TestSuggestions(t *testing.T) {
	 filePath := "../../Dockerfile"
	 result := ReadImage(filePath)
	 if (result != nil) {
		 t.Errorf("Error in reading Dockerfile : %s\n", result)
	 } else {
		 t.Logf("Succesfully ran suggestions in Dockerfile\n")
	 }
}