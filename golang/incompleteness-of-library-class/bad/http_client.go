package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	baseUrl string
}

func NewHttpClient(baseUrl string) *HttpClient {
	return &HttpClient{baseUrl: baseUrl}
}

func (hc HttpClient) Get(endpoint string) (map[string]interface{}, error) {
	url := hc.baseUrl + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	return result, err
}

// Missing POST, PUT, DELETE methods - incomplete library!

func main() {
	client := NewHttpClient("https://jsonplaceholder.typicode.com")

	result, err := client.Get("/posts/1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", result)
}
