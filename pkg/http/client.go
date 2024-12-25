// Package http provides utilities for making HTTP requests and processing
// HTTP responses. It abstracts the standard library's http.Client to simplify
// common HTTP operations.
package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// SendRequest sends an HTTP request to the specified URL using the given method.
// It returns the HTTP headers and body as strings, along with any error encountered.
//
// Parameters:
//   - url: The endpoint to which the request will be sent.
//   - method: The HTTP method to use (e.g., "GET", "POST").
//
// Returns:
//   - headers: A formatted string containing the response headers.
//   - body: The response body as a string.
//   - err: Any error encountered during the request.
//
// Example:
//
//	headers, body, err := SendRequest("https://api.example.com/data", "GET")
//	if err != nil {
//	    log.Fatalf("Request failed: %v", err)
//	}
//	fmt.Println("Headers:", headers)
//	fmt.Println("Body:", body)
func SendRequest(url, method string) (string, string, error) {
	var req *http.Request
	var err error

	if method == "POST" || method == "PUT" {
		reqBody := bytes.NewBuffer([]byte(`{"key":"value"}`))
		req, err = http.NewRequest(method, url, reqBody)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return "", "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	headers := fmt.Sprintf("Status: %s\nHeaders:\n", resp.Status)
	for key, values := range resp.Header {
		for _, value := range values {
			headers += fmt.Sprintf("%s: %s\n", key, value)
		}
	}

	return headers, string(body), nil
}
