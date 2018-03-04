package main

import (
	"encoding/json"
	"testing"
)

func TestRepoIsValid(t *testing.T) {
	validFileContent := `{
		"name": "test_name", 
		"repository": "test.test"
	}`
	invalidNameFileContent := `{
		"name_error": "test_name", 
		"repository": "test.test"
	}`
	invalidRepositoryFileContent := `{
		"name": "test_name", 
		"repository_error": "test.test"
	}`
	invalidBothFileContent := `{
		"name_error": "test_name", 
		"repository_error": "test.test"
	}`

	testCases := []struct {
		scenario string
		input    string
		expected bool
	}{
		{"validFileContent", validFileContent, true},
		{"invalidNameFileContent", invalidNameFileContent, false},
		{"invalidRepositoryFileContent", invalidRepositoryFileContent, false},
		{"invalidBothFileContent", invalidBothFileContent, false},
	}
	for _, testCase := range testCases {
		var repo Repo
		json.Unmarshal([]byte(testCase.input), &repo)
		t.Log("Scenario:", testCase.scenario)
		if repo.isValid() != testCase.expected {
			t.Error(testCase.expected, "!=", repo.isValid())
		}
	}

}
