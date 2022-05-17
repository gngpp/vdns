package vhttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"vdns/lib/util/strs"
)

var defaultClient = NewClient()

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"

// Get Send GET request
func Get(url string, token *string) (response *http.Response, err error) {
	req, err := http.NewRequest(HttpMethodGet.String(), url, nil)
	req.Header.Set("User-Agent", UserAgent)
	if token != nil {
		req.Header.Set("Authorization", "Bearer "+*token)
	}
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
func Post(url string, contentType string, data *interface{}, token *string) (resp *http.Response, err error) {
	var req *http.Request
	if data != nil {
		jsonStr, _ := json.Marshal(data)
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			return nil, err
		}
	}
	req.Header.Set("User-Agent", UserAgent)
	if !strs.IsEmpty(contentType) {
		req.Header.Set("Content-type", contentType)
	}
	if token != nil {
		req.Header.Set("Authorization", "Bearer "+*token)
	}
	if err != nil {
		return nil, err
	}

	resp, err = defaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func IsOK(resp *http.Response) bool {
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusAlreadyReported, http.StatusNoContent:
		return true
	default:
		return false
	}
}
