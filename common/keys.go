package common

// Redis Keys and prefix
var RedisKey = struct {
    LoginWrongTimes    string
    UserToken          string
    ResetPasswordToken string
}{
    LoginWrongTimes:    "LOGIN:WRONG_TIMES:",          // 用户登录失败次数统计 Redis Key 前缀
    UserToken:          "LOGIN:TOKEN:USER:",           // 用户登录 Token Redis Key 前缀
    ResetPasswordToken: "LOGIN:TOKEN:RESET_PASSWORD:", // 用户重置密码 Token 前缀
}
