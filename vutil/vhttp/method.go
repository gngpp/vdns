package vhttp

//goland:noinspection ALL
type HttpMethod string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodTrace   HttpMethod = "TRACE"
	HttpMethodConnect HttpMethod = "CONNECT"
)
