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
    common.DB.Preload("Department").Find(&users)
    response.SuccessWithData(map[string]interface{}{
        "list": users,
    })
}
