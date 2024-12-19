package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	if len(jsonDataList) == 0 {
		return []byte{}, errors.New("list is empty")
	}
	var allStudents []Student
	for i := 0; i < len(jsonDataList); i++ {
		var students []Student
		if err := json.Unmarshal(jsonDataList[i], &students); err != nil {
			return []byte{}, err
		}
		for ii := 0; ii < len(students); ii++ {
			allStudents = append(allStudents, students[ii])
		}
	}
	newJSON, err := json.Marshal(allStudents)
	if err != nil {
		return []byte{}, err
	}
	return newJSON, nil
}

func main() {
	inputJSON1 := []byte(`[
			{
				"name": "Oleg",
				"class": "9B"
			},
			{
				"name": "Ivan",
				"class": "9A"
			}
		]`)

	inputJSON2 := []byte(`[
			{
				"name": "Maria",
				"class": "9B"
			},
			{
				"name": "John",
				"class": "9A"
			}
		]`)

	expectedJSON := []byte(`[{"class":"9B","name":"Oleg"},{"class":"9A","name":"Ivan"},{"class":"9B","name":"Maria"},{"class":"9A","name":"John"}]`)

	mergedJSON, err := mergeJSONData(inputJSON1, inputJSON2)
	if err != nil {
		fmt.Printf("Error while merging JSON data: %v", err)
	}

	if !bytes.Equal(mergedJSON, expectedJSON) {
		fmt.Printf("Expected merged JSON data to be %s, but got %s", expectedJSON, mergedJSON)
	}
}
