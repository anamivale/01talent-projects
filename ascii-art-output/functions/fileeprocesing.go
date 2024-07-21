package functions

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// ValidateBanner checks the file extension, if its found and also if it can be read.
func ValidateBanner(filename string) []string {

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {

		return nil
	}

	if filepath.Ext(filename) != ".txt" {

		return nil
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error in opening file")
		return nil
	}

	defer file.Close()

	var content []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	return content
}
