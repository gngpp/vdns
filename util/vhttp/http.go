package vhttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Get Send GET request
func Get(url string) (response *http.Response, err error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post Send POST request
func Post(url string, data interface{}, contentType string) (response *http.Response, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	response, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
