package scan

import (
	"log"
	"testing"
)

func TestScan(t *testing.T) {
	out, err := Vulnerabilities("nginx")
	if err != nil {
		log.Fatal(err)
	}
	if len(out) < 1 {
		t.Errorf("Error in scanning docker file for vulnerabilities")
	} else {
		t.Logf("Successfully generated report")
	}
}
