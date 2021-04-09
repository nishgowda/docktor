// Package suggestions provides functions to suggest improvements on docker files
// following certain best practices
package suggestions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ReadImage reads a docker file and suggests improvements that can be made
func ReadImage(imagePath string) (string, error) {
	if len(imagePath) < 10  {
		return "", errors.New("File is not a dockerfile")
	}
	// grab the last 10 characters of the filename
	image := imagePath[(len(imagePath) - 10):]
	if runtime.GOOS == "windows" {
		image = strings.TrimRight(image, "\r\n")
	} else {
		image = strings.TrimRight(image, "\n")
	}
	// check if the file is a Dockerfile
	image = strings.ToLower(image)
	if strings.Compare(image, "dockerfile") != 0 {
		return "", errors.New("File given was not a dockerfile")
	}
	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var data DockerVars
	var e ErrorMessages
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	// suggest improvements to each line of the dockerfile
	for scanner.Scan() {
		lineNumber++
		suggestImprovements(scanner.Text(), &data, lineNumber)
	}
	// if certain features are not found in the dockerfile, create the messages
	// and display them to user
	if data.userCount == 0 || data.environCount == 0 || data.workDirCount == 0 {
		createMessages(true, &data, &e)
		displayMessages(&e)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "Detected no issues with Docker container",  nil
}

func createMessages(err bool, data *DockerVars, e *ErrorMessages) error {
	if data == nil || e == nil {
		return errors.New("Invalid arguments")
	}
	if err {
		if data.userCount == 0 {
			e.userMsg = "No user specified in Dockerfile, this is a security risk for your container, consider adding one"
		}
		if data.environCount == 0 {
			e.environMsg = "No environment variable specified in Dockerfile, these are useful for signaling your application is in production."
		}
		if data.workDirCount == 0 {
			e.workDirMsg = "No working directory specified, consider adding one"
		}
		return nil
	}
	return errors.New("No errors detected")
}

func displayMessages(e *ErrorMessages) {
	if e == nil {
		return
	}
	fmt.Println("<---- Detected the following issues with your Dockerfile ---->")
	i := 0
	if len(e.userMsg) != 0 {
		i++
		fmt.Printf("%d.) %s \n", i, e.userMsg)
	}
	if len(e.environMsg) != 0 {
		i++
		fmt.Printf("%d.) %s \n", i, e.environMsg)
	}
	if len(e.workDirMsg) != 0 {
		i++
		fmt.Printf("%d.) %s \n", i, e.workDirMsg)
	}
}

func suggestImprovements(text string, data *DockerVars, lineNumber int) {
	if len(text) == 0 || data == nil || lineNumber < 0 {
		return
	}
	var fromCount int
	words := strings.Fields(text)
	for _, word := range words {
		if fromCount > 0 && !strings.Contains(word, ":") {
			fmt.Printf("Line %d -- Specify a version of this base image: %s\n", lineNumber, word)
		}
		if strings.ToLower(word) == "from" {
			fromCount++
		} else if strings.ToLower(word) == "user" {
			data.userCount++
		} else if strings.ToLower(word) == "env" {
			data.environCount++
		} else if strings.ToLower(word) == "workdir" {
			data.workDirCount++
		} else {
			fromCount = 0
		}
	}
}
