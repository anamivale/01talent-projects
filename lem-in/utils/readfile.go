package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Failed to read the file data.")
		return nil
	}
	content := string(file)
	slice := strings.Split(content, "\n")
	return slice
}
