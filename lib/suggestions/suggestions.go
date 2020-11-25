package suggestions

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"errors"
)

// ReadImage reads a docker file and ouptuts its contents
func ReadImage(imagePath string) error {
	image := imagePath[(len(imagePath) - 10): len(imagePath)]
	if runtime.GOOS == "windows" {
		image = strings.TrimRight(image, "\r\n")
	} else {
		image = strings.TrimRight(image, "\n")
	}
	
	image = strings.ToLower(image)
	if strings.Compare(image, "dockerfile") != 0 {
		return errors.New("File given was not a dockerfile")
	}
	file, err := os.Open(imagePath)
	if err != nil {
		return errors.Unwrap(err)
	}
	defer file.Close()

	var data DockerVars
	var e ErrorMessages
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumber ++
		suggestImprovements(scanner.Text(), &data, lineNumber)
	}
	if data.userCount == 0 || data.environCount == 0 || data.workDirCount == 0 {
		createMessages(true, &data, &e)
		displayMessages(&e)
	}
	if err := scanner.Err(); err != nil {
		return errors.Unwrap(err)
	}
	return nil
}

func createMessages(err bool, data *DockerVars, e *ErrorMessages) error {
	if err {
		if data.userCount == 0 {
			e.userMsg = "No user specified in Dockerfile, this is a security risk for your container, consider adding one"
		} 
		if data.environCount == 0 {
			e.environMsg = "No environment variable specified in Dockerfile, consider adding one as it helps secure your repository and is considered a best practice"
		}  
		if data.workDirCount == 0 {
			e.workDirMsg = "No working directory specified, consider adding one"
		}
		return nil
	} 
	return errors.New("No errors detected")
}

func displayMessages(e *ErrorMessages) {
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
	var fromCount int
	words := strings.Fields(text)
	for _, word := range words {
		if fromCount > 0  {
			if !strings.Contains(word, ":") {
				fmt.Printf("Line %d -- Specifiy a version of this base image: %s\n", lineNumber, word)
			}
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

