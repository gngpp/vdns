package vhttp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var defaultClient = CreateClient()

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"

// Get Send GET request
func Get(url string) (response *http.Response, err error) {
	req, err := http.NewRequest(HttpMethodGet.String(), url, nil)
	req.Header.Set("User-Agent", UserAgent)
	if err != nil {
		return nil, err
	}
	resp, err := defaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post Send POST request
func Post(url string, contentType string, data interface{}) (response *http.Response, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Add("Content-type", contentType)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	response, err = defaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func IsOK(resp *http.Response) bool {
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusAccepted:
	case http.StatusAlreadyReported:
	case http.StatusNoContent:
		return true
	default:
		return false
	}
	return false
}
