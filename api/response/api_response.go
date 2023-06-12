package response

type APICreateResponse struct {
	APIDescription string `json:"api_description"`
	APIEndpoint    string `json:"api_endpoint"`
	RequestMethod  string `json:"request_method"`
	IsOpen         bool   `json:"is_open"`
	CreatedAt      string `json:"created_at"`
}

type APIGetResponse struct {
	ID             uint   `json:"id"`
	APIDescription string `json:"api_description"`
	APIEndpoint    string `json:"api_endpoint"`
	RequestMethod  string `json:"request_method"`
	IsOpen         bool   `json:"is_open"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
