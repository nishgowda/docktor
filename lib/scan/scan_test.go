package scan

import (
	"testing"
)

func TestScan(t *testing.T) {
	out := Vulnerabilities("nginx")
	if len(out) < 1 {
		t.Errorf("Error in scanning docker file for vulnerabilities")
	} else {
		t.Logf("Succesfully generated report")
	}
}