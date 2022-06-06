package cloudflare_model

// UpdateRecordResponse represents the response from the DNS endpoint.
type UpdateRecordResponse struct {
	Response
	Result Result `json:"result"`
}
