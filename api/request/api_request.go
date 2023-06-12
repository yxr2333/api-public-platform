package request

type APICreateRequest struct {
	APIDescription string `json:"api_description" binding:"required"`
	APIEndpoint    string `json:"api_endpoint" binding:"required"`
	RequestMethod  string `json:"request_method" binding:"required"`
	IsOpen         *bool  `json:"is_open"`
}

type APIUpdateRequest struct {
	ID             uint   `json:"id" binding:"required"`
	APIDescription string `json:"api_description" binding:"required"`
	APIEndpoint    string `json:"api_endpoint" binding:"required"`
	RequestMethod  string `json:"request_method" binding:"required"`
	IsOpen         *bool  `json:"is_open"`
}
