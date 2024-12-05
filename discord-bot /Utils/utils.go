package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// MakeHTTPGetRequest makes a GET request to a given URL and returns the response body
func MakeHTTPGetRequest(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add optional headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

// ParseJSON decodes a JSON response into a given struct
func ParseJSON(data []byte, target interface{}) error {
	err := json.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}
	return nil
}

// LogError logs an error message if err is not nil
func LogError(message string, err error) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}

// FormatURL formats a base URL with query parameters
func FormatURL(baseURL string, params map[string]string) string {
	url := baseURL + "?"
	for key, value := range params {
		url += fmt.Sprintf("%s=%s&", key, value)
	}
	return url[:len(url)-1] // Remove the trailing "&"
}

