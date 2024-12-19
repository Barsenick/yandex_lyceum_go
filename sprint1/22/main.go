package main

import (
	"encoding/json"
	"errors"
)

type Student struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func has_key(m map[string][]byte, v string) bool {
	for i := range m {
		if i == v {
			return true
		}
	}
	return false
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	if len(jsonData) == 0 {
		return map[string][]byte{}, errors.New("list is empty")
	}
	var allStudents []Student
	if err := json.Unmarshal(jsonData, &allStudents); err != nil {
		return map[string][]byte{}, err
	}
	m := make(map[string][]byte, 0)
	for i := 0; i < len(allStudents); i++ {
		if has_key(m, allStudents[i].Class) {
			continue
		}
		classStudents := make([]Student, 0)
		class := allStudents[i].Class
		for ii := 0; ii < len(allStudents); ii++ {
			if allStudents[ii].Class == class {
				classStudents = append(classStudents, allStudents[ii])
			}
		}
		classJSON, err := json.Marshal(classStudents)
		if err != nil {
			return map[string][]byte{}, err
		}
		m[class] = classJSON
	}
	return m, nil
}
