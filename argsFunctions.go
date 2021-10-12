package main

import (
	"errors"
	"math"
	"os"
)

func getFilePrefix() (string, error) {
	if len(os.Args) < 3 {
		return "", errors.New("Wrong format of variables. May be you forgot to name some variable or forgot to give the value.")
	}

	return os.Args[2], nil
}

func getArgsVariables() (map[string]string, error) {
	variablesLen := float64(len(os.Args) - 3)
	if variablesLen == 0 || math.Mod(variablesLen, 2) != 0 {
		return nil, errors.New("Wrong format of variables. May be you forgot to name some variable or forgot to give the value.")
	}

	variables := make(map[string]string)

	for i := 3; i < len(os.Args); i += 2 {
		arg := os.Args[i]

		if len(arg) > 2 && arg[0:2] == "--" {
			variables[arg[2:]] = os.Args[i+1]
		}
	}

	return variables, nil
}
