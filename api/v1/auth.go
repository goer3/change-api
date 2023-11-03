package v1

import (
    "change-api/common"
    "change-api/dto"
    "change-api/model"
    "change-api/pkg/gedis"
    "change-api/pkg/response"
    "change-api/pkg/utils"
    "github.com/gin-gonic/gin"
)

// 用户激活
func ActiveUserHandler(ctx *gin.Context) {
    // 获取 URL 参数
    token := ctx.Param("token")
    if token == "" {
        response.FailedWithMessage("获取用户重置密码 Token 失败")
    }

    // 根据 Token 查询重置密码用户
    conn := gedis.NewOperation()
    key := common.RedisKey.ResetPasswordToken + token
    id := conn.GetString(key).Unwrap()
    if id == "" {
        response.FailedWithMessage("重置密码 Token 已失效，请重新获取")
    }

    // 验证是否获取到密码
    var req dto.ResetPassword
    if err := ctx.ShouldBind(&req); err != nil {
        response.FailedWithMessage("获取提交的密码数据失败")
    }

    // 验证密码复杂度
    if utils.PasswordComplexity(req.Password) < common.Config.Login.PasswordLevel {
        response.FailedWithMessage(common.PasswordLevelMessage[common.Config.Login.PasswordLevel])
    }

    // 查询用户，修改密码
    var user model.User
    err := common.DB.Where("id = ?", id).First(&user).Error
    if err != nil {
        response.FailedWithMessage("获取修改密码的用户信息失败")
    }

    // 修改密码和用户状态
    if err = common.DB.Model(&user).Updates(map[string]interface{}{
        "password": utils.CryptoPassword(req.Password),
        "status":   1,
    }).Error; err != nil {
        response.FailedWithMessage("更新用户密码和用户状态失败")
    }

    // 成功相应
    response.Success()
}
