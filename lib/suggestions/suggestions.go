package suggestions

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
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
		fmt.Println("File given was not a dockerfile")
		os.Exit(1)
	}
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var data DockerVars
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumber ++
		suggestImprovements(scanner.Text(), &data, lineNumber)
	}
	if data.userCount == 0 {
		fmt.Println("No user specified in Dockerfile, this is a security risk for your container, consider adding one")
	} else if data.environCount == 0 {
		fmt.Println("No environment variable specified in Dockerfile, consider adding one as it helps secure your repository and is considered a best practice")
	} else if data.environCount == 0 {
		fmt.Println("No working direcotry specified, consider adding one as")
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
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

