package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"text/template"
)

func isValidPath(path string) bool {
	if path[0] == '.' && path[1] == '/' && path[len(path)-1] != '/' && len(path) > 2 {
		return true
	}
	return false
}

func getVariables() (map[string]string, error) {
	variablesLen := float64(len(os.Args) - 2)
	if variablesLen == 0 || math.Mod(variablesLen, 2) != 0 {
		return nil, errors.New("Wrong format of variables. May be you forgot to name some variable or forgot to give the value.")
	}

	variables := make(map[string]string)

	for i := 0; i < len(os.Args); i += 2 {
		arg := os.Args[i]
		if len(arg) > 2 && arg[0:2] == "--" {
			variables[arg[2:]] = os.Args[i+1]
		}
	}

	return variables, nil
}

func createPath(path string) (string, error) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return currentDirectory + path[1:], nil
}

func readFileContent(path string) (string, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", nil
	}
	return string(fileContent), nil
}

func genFromTemplate(temaplateName string, content string, variables map[string]string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	t := template.Must(template.New(temaplateName).Parse(content))
	err := t.Execute(buf, variables)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func fileTargetPath(path string) (string, error) {
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
	return currentDirectory + path[lastBar:], nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please write the relative path to the source file or to the directory.")
		return
	}

	path := os.Args[1]

	valid := isValidPath(path)
	if !valid {
		fmt.Println("Invalid source path\n\tThe format is: ./GoCodeGen ./<my-relative-path>")
		return
	}

	variables, err := getVariables()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	globalPath, err := createPath(path)
	if err != nil {
		fmt.Println("Coundn't get the current path of the program execution to read the file")
		return
	}

	content, err := readFileContent(globalPath)
	if err != nil {
		fmt.Println("Couldn't read the content of file with path: " + globalPath)
		return
	}

	resultBuf, err := genFromTemplate("test1", content, variables)
	if err != nil {
		fmt.Println("Couldn't generate from file wih path " + globalPath)
		return
	}

	targetPath, err := fileTargetPath(path)
	if err != nil {
		fmt.Println("Couldn't write file wih path " + globalPath)
		return
	}
	os.WriteFile(targetPath, resultBuf.Bytes(), 0644)
}
