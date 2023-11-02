package v1

import (
    "change-api/pkg/response"
    "github.com/gin-gonic/gin"
)

// 用户激活
func ActiveUserHandler(ctx *gin.Context) {
    response.Success()
}
