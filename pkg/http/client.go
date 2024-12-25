package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// SendRequest handles making an HTTP request
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
