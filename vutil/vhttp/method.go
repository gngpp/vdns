package vhttp

//goland:noinspection ALL
type HttpMethod string

type HttpMethodEnum struct {
	HttpMethodGet     HttpMethod
	HttpMethodPut     HttpMethod
	HttpMethodPost    HttpMethod
	HttpMethodDelete  HttpMethod
	HttpMethodHead    HttpMethod
	HttpMethodOptions HttpMethod
	HttpMethodTrace   HttpMethod
	HttpMethodConnect HttpMethod
}

//goland:noinspection ALL
var HttpMethods = HttpMethodEnum{
	HttpMethodGet:     "GET",
	HttpMethodPut:     "PUT",
	HttpMethodHead:    "HEAD",
	HttpMethodOptions: "OPTIONS",
	HttpMethodDelete:  "DELETE",
	HttpMethodPost:    "POST",
	HttpMethodTrace:   "TRACE",
	HttpMethodConnect: "CONNECT",
}
