package initialize

import (
    "change-api/common"
    "change-api/middleware"
    "change-api/pkg/log2"
    "change-api/routes"

    "github.com/gin-gonic/gin"
)

// 路由初始化
func Router() *gin.Engine {
    // 创建一个没中间件的路由引擎
    r := gin.New()

    // 中间件
    r.Use(middleware.AccessLog) // 请求日志中间件
    r.Use(middleware.Cors)      // 跨域配置中间件
    r.Use(middleware.Exception) // 异常拦截中间件

    // 路由组
    rg := r.Group(common.Config.System.ApiPrefix + "/" + common.Config.System.ApiVersion)
    {
        routes.PublicRoute(rg) // 开放路由组
    }

    // 初始化成功
    log2.INFO("router initialize success")
    return r
}
