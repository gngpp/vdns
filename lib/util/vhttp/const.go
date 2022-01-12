package vhttp

//goland:noinspection ALL
type HttpMethod string

func (receiver HttpMethod) String() string {
	return string(receiver)
}

type HttpHeader string

func (receiver HttpHeader) String() string {
	return string(receiver)
}

// methods
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

// headers
const (
	Authorization HttpHeader = "Authorization"
	ContentType   HttpHeader = "Content-Type"
)

// header value

const (
	ApplicationJson = "application/json"
)
