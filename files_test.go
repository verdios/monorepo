package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"test.txt", false},
		{"./test.txt", false},
		{"files_test.go", true},
		{"./files_test.go", true},
	}
	for _, testCase := range testCases {
		fileExists := fileExists(testCase.input)
		if fileExists != testCase.expected {
			t.Error(fileExists, "!=", testCase.expected)
		}
	}
}

func TestCreateConfigFile(t *testing.T) {
	if err := createConfigFile("not_found"); err == nil {
		t.Error("Should be throw an error at this point")
	}
	err := createConfigFile("files_test.go")
	if err != nil {
		t.Error(err)
	}
	if !fileExists(monorepoFilename) {
		t.Error("Not found:", monorepoFilename)
	}
	if fileExists(monorepoFilename) {
		if err := os.Remove(monorepoFilename); err != nil {
			t.Fatal(err)
		}
	}
}

func TestReadConfigFileSucess(t *testing.T) {
	filename := "test_file.json"
	fileContent := `{
		"repositories": [
			{
				"name": "test_name", 
				"repository": "test.test"
			}
		]
	}`
	if err := ioutil.WriteFile(filename, []byte(fileContent), filePermission); err != nil {
		t.Fatal(err)
	}
	if err := createConfigFile(filename); err != nil {
		t.Fatal(err)
	}
	repositories, err := readConfigFile()
	if err != nil {
		t.Error(err)
	}
	if len(repositories.Repos) > 1 {
		t.Error("Number of repositories should be 1:", len(repositories.Repos))
	}
	if repositories.Repos[0].Name != "test_name" || repositories.Repos[0].Repository != "test.test" {
		t.Error("test_name test.test !=", repositories.Repos[0].Name, repositories.Repos[0].Repository)
	}
	if fileExists(filename) {
		if err := os.Remove(filename); err != nil {
			t.Fatal(err)
		}
	}
	if fileExists(monorepoFilename) {
		if err := os.Remove(monorepoFilename); err != nil {
			t.Fatal(err)
		}
	}
}

func TestReadConfigFileError(t *testing.T) {
	filename := "test_file.json"
	fileContent := `{
		"repositories": [
			{
				"name": "test_name", 
				"repository": "test.test"
			
		]
	}`
	if err := ioutil.WriteFile(filename, []byte(fileContent), filePermission); err != nil {
		t.Fatal(err)
	}
	if err := createConfigFile(filename); err != nil {
		t.Fatal(err)
	}
	_, err := readConfigFile()
	if err == nil {
		t.Error("Should be throw an error at this point")
	}
	if fileExists(filename) {
		if err := os.Remove(filename); err != nil {
			t.Fatal(err)
		}
	}
	if fileExists(monorepoFilename) {
		if err := os.Remove(monorepoFilename); err != nil {
			t.Fatal(err)
		}
	}
}

func TestReadConfigNoConfigFile(t *testing.T) {
	filename := "test_file.json"
	fileContent := `{
		"repositories": [
			{
				"name": "test_name", 
				"repository": "test.test"
			}
		]
	}`
	if err := ioutil.WriteFile(filename, []byte(fileContent), filePermission); err != nil {
		t.Fatal(err)
	}
	if err := createConfigFile(filename); err != nil {
		t.Fatal(err)
	}
	if fileExists(filename) {
		if err := os.Remove(filename); err != nil {
			t.Fatal(err)
		}
	}
	if fileExists(monorepoFilename) {
		if err := os.Remove(monorepoFilename); err != nil {
			t.Fatal(err)
		}
	}
	_, err := readConfigFile()
	if err == nil {
		t.Error("Should be throw an error at this point")
	}
}
func TestGetInput(t *testing.T) {
	testCases := []struct {
		input        []string
		fileExpected string
		errExpected  error
	}{
		{[]string{}, "", noInputFileError{}},
		{[]string{"test", "test"}, "", fileNotFoundError{}},
		{[]string{"test", "files_test.go"}, "files_test.go", nil},
	}

	for _, testCase := range testCases {
		file, err := getInputFile(testCase.input)
		if file != testCase.fileExpected {
			t.Error(file, "!=", testCase.fileExpected)
		}
		if testCase.errExpected != nil {
			if err != testCase.errExpected {
				t.Error(err, "!=", testCase.errExpected)
			}
			if err.Error() != testCase.errExpected.Error() {
				t.Error(err.Error(), "!=", testCase.errExpected.Error())
			}
		}
	}
}
