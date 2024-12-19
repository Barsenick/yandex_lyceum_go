package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student
	if err := json.Unmarshal(jsonData, &students); err != nil {
		return []byte{}, err
	}
	for i := 0; i < len(students); i++ {
		students[i].Grade += 1
	}
	newJSON, err := json.Marshal(students)
	if err != nil {
		return []byte{}, err
	}
	return newJSON, nil
}

func main() {
	inputJSON := []byte(`[
			{
				"name": "Oleg",
				"grade": 12
			}
		]`)
	expectedJSON := []byte(`[{"name":"Oleg","grade":13}]`)
	newJSON, err := modifyJSON(inputJSON)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		if bytes.Equal(expectedJSON, newJSON) {
			fmt.Printf("Success: %v", string(newJSON))
		} else {
			fmt.Printf("Expected '%v', got '%v'", expectedJSON, newJSON)
		}
	}
}
