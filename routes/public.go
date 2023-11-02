package routes

import (
    "change-api/api"
    v1 "change-api/api/v1"
    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
)

// 开放路由组
func PublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
    rg.GET("/health", api.HealthHandler)           // 健康检查接口
    rg.GET("/info", api.InfoHandler)               // 系统信息接口
    rg.GET("/version", api.VersionHandler)         // 版本信息接口
    rg.POST("/login", auth.LoginHandler)           // 用户登录
    rg.PUT("/active/:token", v1.ActiveUserHandler) // 用户未激活，登录重置密码
    return rg
}
