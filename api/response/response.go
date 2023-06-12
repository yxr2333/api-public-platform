package response

type BaseResponse struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
