package dto

// 用户登录请求
type LoginRequest struct {
    Account  string `json:"account"`
    Password string `json:"password"`
}

// 用户登录响应
type LoginResponse struct {
    Token  string `json:"token"`
    Expire string `json:"expire"`
}

// 重置密码
type ResetPassword struct {
    Password   string `json:"password"`
    RePassword string `json:"re_password"`
}
