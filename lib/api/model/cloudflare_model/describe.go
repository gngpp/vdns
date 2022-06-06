package cloudflare_model

type DescribeRecordResponse struct {
	Response
	Result     []Result   `json:"result,omitempty"`
	ResultInfo ResultInfo `json:"result_info,omitempty"`
}

type ResultInfo struct {
	Page       *int64 `json:"page,omitempty"`
	PerPage    *int64 `json:"per_page,omitempty"`
	Count      *int64 `json:"count,omitempty"`
	TotalCount *int64 `json:"total_count,omitempty"`
	TotalPages *int64 `json:"total_pages,omitempty"`
}
