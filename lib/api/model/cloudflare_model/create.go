package cloudflare_model

// CreateRecordResponse represents the response from the DNS endpoint.
type CreateRecordResponse struct {
	Response
	Result Result `json:"result"`
}
