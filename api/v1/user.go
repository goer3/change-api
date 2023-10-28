package v1

import (
    "change-api/common"
    "change-api/model"
    "change-api/pkg/response"
    "github.com/gin-gonic/gin"
)

// 用户列表
func UserListHandler(ctx *gin.Context) {
    var users []model.User
    common.DB.Find(&users)
    response.SuccessWithData(map[string]interface{}{
        "list": users,
    })
}

// 删除用户
func UserDeleteHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    common.DB.Where("id = ?", id).Delete(&model.User{})
    response.Success()
}
