// Package scan scans a docker image for vulnerabilities using docker scan
package scan

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Vulnerabilities scans images for vulnerabilities
func Vulnerabilities(image string) (string, error) {
	app := "docker"
	arg1 := "scan"
	// group issues in json format
	arg2 := "--json"
	arg3 := "--group-issues"
	if len(image) < 1 {
		return "", errors.New("No image specified")
	}
	cmd := exec.Command(app, arg1, arg2, arg3, image)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), nil
	}
	// checks if Snyk is enabeld
	if strings.Compare(string(output), "Docker Scan relies upon access to Snyk, a third party provider, do you consent to proceed using Snyk? (y/N)") == 0 {
		return "", errors.New("You need to enable Snyk to use this feature")
	}
	return string(output), nil
}

// WriteFile writes output to file
func WriteFile(text string, filename string) (string, error) {
	if len(filename) < 1 || len(text) < 1 {
		return "", errors.New("No filename or text provided")
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, text)
	w.Flush()
	msg := "Successfully wrote vulnerability report to " + filename
	return msg, nil
}
