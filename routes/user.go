package routes

import (
    v1 "change-api/api/v1"
    "github.com/gin-gonic/gin"
)

// 用户路由组
func UserRoutes(rg *gin.RouterGroup) gin.IRoutes {
    rg.GET("/list", v1.UserListHandler)     // 用户列表接口
    rg.DELETE("/:id", v1.UserDeleteHandler) // 用户删除接口
    return rg
}
