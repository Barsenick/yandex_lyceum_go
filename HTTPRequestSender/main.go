package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8080/api/v1/calculate"

	requestBody := []byte(
		`{"expression":"51+1"}`)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	log.Println("Response Status:", response.Status)
	log.Println("Response Body:", string(bodyBytes))
}
