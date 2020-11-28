// Package scan scans a docker image for vulnerabilities using docker scan
package scan

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Vulnerabilities scans images for vulnerabilities
func Vulnerabilities(image string) string {
	app := "docker"
	arg := "scan"
	if len(image) < 1 {
		log.Fatal("No image specified")
		return ""
	}
	cmd := exec.Command(app, arg, image)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output)
	}
	return string(output)
}

// WriteFile writes output to file
func WriteFile(text string, filename string) error {
	if len(filename) < 1 || len(text) < 1 {
		return errors.New("No filename or text provided")
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, text)
	w.Flush()
	fmt.Printf("Successfully wrote vulnerability report to %s\n", filename)
	return nil
}
