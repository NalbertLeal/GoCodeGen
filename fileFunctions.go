package main

import (
	"os"
)

func readFileContent(path string) (string, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", nil
	}
	return string(fileContent), nil
}
