package dto

// 用户登录
type LoginRequest struct {
    Account  string `json:"account"`
    Password string `json:"password"`
}
