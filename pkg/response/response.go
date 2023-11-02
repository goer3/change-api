package response

// 用户状态码
const (
    OK           = 200
    NotOK        = 400
    Forbidden    = 403
    NotFound     = 404
    ParamError   = 406
    ServerError  = 500
    Unauthorized = 1000
    UnActived    = 1001
)

// 用户状态码对应的错误信息
const (
    OKMessage           = "操作成功"
    NotOKMessage        = "操作失败"
    ForbiddenMessage    = "无权限访问该资源"
    NotFoundMessage     = "资源未找到"
    ParamErrorMessage   = "参数错误"
    ServerErrorMessage  = "服务器错误，请联系管理员"
    UnauthorizedMessage = "用户未登录"
    UnActivedMessage    = "用户未激活"
)

// 用户状态码和消息绑定
var CustomMessage = map[int]string{
    OK:           OKMessage,
    NotOK:        NotOKMessage,
    Forbidden:    ForbiddenMessage,
    NotFound:     NotFoundMessage,
    ParamError:   ParamErrorMessage,
    ServerError:  ServerErrorMessage,
    Unauthorized: UnauthorizedMessage,
    UnActived:    UnActivedMessage,
}
