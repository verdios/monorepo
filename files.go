package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	monorepoFilename = ".monorepo"
	filePermission   = 0644
)

func getInputFile(args []string) (string, error) {
	if len(args) <= 0 {
		return "", noInputFileError{}
	}
	if fileExists(args[1]) {
		return args[1], nil
	}
	return "", fileNotFoundError{}
}

func createConfigFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(monorepoFilename, data, filePermission)
}

func readConfigFile() (Repositories, error) {
	var repositories Repositories
	data, err := ioutil.ReadFile(monorepoFilename)
	if err != nil {
		return repositories, err
	}
	if err := json.Unmarshal(data, &repositories); err != nil {
		return repositories, err
	}
	return repositories, nil
}

func fileExists(filename string) bool {
	if !strings.Contains(filename, "/") {
		dir, _ := os.Getwd()
		filename = fmt.Sprintf("%s/%s", dir, filename)
	}
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
