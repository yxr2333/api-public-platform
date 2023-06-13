package response

type UserRegisterResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
}

type UserLoginResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

type UserBaseInfo struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
}
