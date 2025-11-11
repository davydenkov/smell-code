package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	baseUrl        string
	defaultHeaders map[string]string
}

func NewHttpClient(baseUrl string, defaultHeaders map[string]string) *HttpClient {
	if defaultHeaders == nil {
		defaultHeaders = make(map[string]string)
	}

	// Set default content type if not provided
	if _, exists := defaultHeaders["Content-Type"]; !exists {
		defaultHeaders["Content-Type"] = "application/json"
	}

	return &HttpClient{
		baseUrl:        baseUrl,
		defaultHeaders: defaultHeaders,
	}
}

func (hc HttpClient) Get(endpoint string, headers map[string]string) (map[string]interface{}, error) {
	return hc.request("GET", endpoint, nil, headers)
}

func (hc HttpClient) Post(endpoint string, data interface{}, headers map[string]string) (map[string]interface{}, error) {
	return hc.request("POST", endpoint, data, headers)
}

func (hc HttpClient) Put(endpoint string, data interface{}, headers map[string]string) (map[string]interface{}, error) {
	return hc.request("PUT", endpoint, data, headers)
}

func (hc HttpClient) Delete(endpoint string, headers map[string]string) (map[string]interface{}, error) {
	return hc.request("DELETE", endpoint, nil, headers)
}

func (hc HttpClient) Patch(endpoint string, data interface{}, headers map[string]string) (map[string]interface{}, error) {
	return hc.request("PATCH", endpoint, data, headers)
}

func (hc HttpClient) request(method, endpoint string, data interface{}, additionalHeaders map[string]string) (map[string]interface{}, error) {
	url := hc.baseUrl + endpoint

	// Merge headers
	headers := make(map[string]string)
	for k, v := range hc.defaultHeaders {
		headers[k] = v
	}
	for k, v := range additionalHeaders {
		headers[k] = v
	}

	var body io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	return result, err
}

func main() {
	client := NewHttpClient("https://jsonplaceholder.typicode.com", nil)

	// Test GET
	result, err := client.Get("/posts/1", nil)
	if err != nil {
		fmt.Printf("GET Error: %v\n", err)
		return
	}
	fmt.Printf("GET Response: %+v\n", result)

	// Test POST
	postData := map[string]interface{}{
		"title":  "Test Post",
		"body":   "This is a test",
		"userId": 1,
	}
	postResult, err := client.Post("/posts", postData, nil)
	if err != nil {
		fmt.Printf("POST Error: %v\n", err)
		return
	}
	fmt.Printf("POST Response: %+v\n", postResult)
}
