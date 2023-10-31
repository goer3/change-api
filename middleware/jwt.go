package middleware

import (
    "change-api/common"
    "change-api/dto"
    "change-api/model"
    "change-api/pkg/gedis"
    "change-api/pkg/utils"
    "errors"
    "time"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
)

// JWT 认证中间件
func JWTAuth() (*jwt.GinJWTMiddleware, error) {
    return jwt.New(&jwt.GinJWTMiddleware{
        Realm:           common.Config.JWT.Realm,                                // JWT 标识
        Key:             []byte(common.Config.JWT.Key),                          // 签名 Key
        Timeout:         time.Duration(common.Config.JWT.Timeout) * time.Second, // Token 有效期
        Authenticator:   authenticator,                                          // 用户登录校验
        PayloadFunc:     payloadFunc,                                            // Token 封装
        LoginResponse:   loginResponse,                                          // 登录成功响应
        Unauthorized:    unauthorized,                                           // 登录，认证失败响应
        IdentityHandler: identityHandler,                                        // 解析 Token
        Authorizator:    authorizator,                                           // 验证 Token
        LogoutResponse:  logoutResponse,                                         // 注销登录
        TokenLookup:     "header: Authorization, query: token, cookie: jwt",     // Token 查找的字段
        TokenHeadName:   "Bearer",                                               // Token 请求头名称
    })
}

// 隶属 Login 中间件，当调用 LoginHandler 就会触发
// 通过从 ctx 中检索出数据，进行用户登录认证
// 返回包含用户信息的 Map 或者 Struct
func authenticator(ctx *gin.Context) (interface{}, error) {
    // 获取用户登录数据
    var req dto.LoginRequest
    if err := ctx.ShouldBind(&req); err != nil {
        return nil, errors.New("获取用户登录信息失败")
    }

    // 通过 IP 判断用户登录次数是否达到上限
    // 1.获取客户端 IP，注意确保获取到的是真实 IP
    ip := ctx.ClientIP()
    if ip == "" {
        ip = "None"
    }

    // 2.获取 redis 中该用户登录错误次数
    key := req.Account + ":" + ip
    var conn = gedis.NewOperation()
    times := conn.GetInt(key).UnwrapWithDefault(0)
    // 密码错误次数，默认允许 5 次
    if times >= 5 {

    }

    // 查询模板
    db := common.DB
    var user model.User
    var err error

    // 判断用户登录采用的方式，支持工号，手机号，Email
    if utils.IsMobile(req.Account) {
        err = db.Where("mobile = ?", req.Account).First(&user).Error
    } else if utils.IsEmail(req.Account) {
        err = db.Where("email = ?", req.Account).First(&user).Error
    } else {
        err = db.Where("job_id = ?", req.Account).First(&user).Error
    }

    // 判断查询结果
    if err != nil {
        common.Log.Error(err)
        return nil, errors.New("用户名或密码错误")
    }

    // 密码校验
    if !utils.ComparePassword(user.Password, req.Password) {
        //
        return nil, errors.New("用户名或密码错误")
    }

    // 用户状态校验
    switch user.Status {
    case common.BrokenStatus:
        return nil, errors.New("用户已禁用，请联系管理员")
    case common.UnActiveStatus:
        return nil, errors.New("用户未激活，跳转到激活页面，修改密码激活")
    case common.LockedStatus:
        return nil, errors.New("用户已锁定，请等待多长时间后重新登录")
    }

    // 位置状态
    if user.Status != common.NormalStatus {
        return nil, errors.New("用户处于未知状态，请联系管理员")
    }

}
