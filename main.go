package main

import (
	"fmt"
	"os"
)

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

	variables, err := getArgsVariables()
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

	generationName, err := getFilePrefix()
	if err != nil {
		fmt.Println("Couldn't get the generation name from args")
		return
	}

	targetPath, err := fileTargetPath(generationName, path)
	if err != nil {
		fmt.Println("Couldn't write file with path " + globalPath)
		return
	}
	os.WriteFile(targetPath, resultBuf.Bytes(), 0644)
}
