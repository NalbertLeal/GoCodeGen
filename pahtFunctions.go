package main

import (
	"os"
)

func isValidPath(path string) bool {
	if path[0] == '.' && path[1] == '/' && path[len(path)-1] != '/' && len(path) > 2 {
		return true
	}
	return false
}

func createPath(path string) (string, error) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return currentDirectory + path[1:], nil
}

func fileTargetPath(generationName string, path string) (string, error) {
	lastBar := -1
	for i := len(path) - 1; i > -1; i-- {
		if path[i] == '/' {
			lastBar = i
			break
		}
	}

	currentDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return currentDirectory + "/" + generationName + "_" + path[lastBar+1:], nil
}
